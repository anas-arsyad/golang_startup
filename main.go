package main

import (
	"bwastartup/auth"
	"bwastartup/handler"
	"bwastartup/helper"
	"bwastartup/user"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	/* START */
	username := helper.EnvVariable("USERNAME_DB")
	password := helper.EnvVariable("PASSWORD_DB")
	port := helper.EnvVariable("PORT_DB")
	table := helper.EnvVariable("TABLE_DB")
	/* END */

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, port, table)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("DATABASE OK")

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	authService := auth.NewServiceJwt()

	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()
	router.SetTrustedProxies([]string{"localhost"})
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "OK")
	})
	api := router.Group("/api/v1")
	user := api.Group("/user")

	user.POST("/register", userHandler.RegisterUser)
	user.POST("/login", userHandler.LoginUser)
	user.POST("/email_checker", userHandler.IsEmailAvailable)
	user.POST("/avatars", userHandler.UploadAvatar)

	router.Run("localhost:8080")

}

/*
input
handler
service
repository
db

*/
