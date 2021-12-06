package main

import (
	"bwastartup/auth"
	"bwastartup/campaign"
	"bwastartup/handler"
	"bwastartup/helper"
	"bwastartup/transaction"
	"bwastartup/user"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	/* START */
	username := helper.EnvVariable("USERNAME_DB")
	password := helper.EnvVariable("PASSWORD_DB")
	port := helper.EnvVariable("PORT_DB")
	table := helper.EnvVariable("TABLE_DB")
	/* END */

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, port, table)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("DATABASE OK")

	userRepository := user.NewRepository(db)
	campaignRepository := campaign.NewRepository(db)
	transactioRepository := transaction.NewRepository(db)

	authService := auth.NewServiceJwt()
	userService := user.NewService(userRepository)
	campaignService := campaign.NewService(campaignRepository)
	transactionService := transaction.NewService(transactioRepository)

	// a, _ := campaignService.GetCampaignById(1)
	// toLog, _ := json.MarshalIndent(a, "", "  ")
	// fmt.Println(string(toLog))

	userHandler := handler.NewUserHandler(userService, authService)
	campaignHandler := handler.NewCampaignHandler(campaignService)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	router := gin.Default()
	router.Static("/images", "./images")
	router.SetTrustedProxies([]string{"localhost"})
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "OK")
	})

	api := router.Group("/api/v1")

	/* USER */
	user := api.Group("/user")
	user.POST("/register", userHandler.RegisterUser)
	user.POST("/login", userHandler.LoginUser)
	user.POST("/email_checker", userHandler.IsEmailAvailable)
	user.POST("/avatars", authMiddleware(authService, userService), userHandler.UploadAvatar)

	/* CAMPAIGN */
	campaign := api.Group("/campaign")
	campaign.GET("/", campaignHandler.GetCampaign)
	campaign.GET("/:id", campaignHandler.GetCampaignById)
	campaign.POST("/", authMiddleware(authService, userService), campaignHandler.CreateCampaign)
	campaign.PUT("/:id", authMiddleware(authService, userService), campaignHandler.UpdateCampaign)
	campaign.POST("/campaign-images", authMiddleware(authService, userService), campaignHandler.UploadImage)

	/* TRANSACTION */
	// transaction:=api.Group("/transaction")
	campaign.GET("/:id/transactions", transactionHandler.GetCampaignById)

	router.Run("localhost:8080")

}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			response := helper.ApiResponse("Unauthorized", http.StatusUnauthorized, "fail", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		arrayToken := strings.Split(authHeader, " ")
		tokenString := ""
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.ApiResponse("Unauthorized", http.StatusUnauthorized, "fail", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.ApiResponse("Unauthorized", http.StatusUnauthorized, "fail", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		userId := int(claim["user_id"].(float64))
		user, err := userService.GetUserById(userId)
		if err != nil {
			response := helper.ApiResponse("Unauthorized", http.StatusUnauthorized, "fail", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		c.Set("currentUser", user)
	}

}

/*
input
handler
service
repository
db

*/

/*
--== Auth ==--
ambil nilai header authorization :Bearer
dari header authorization kita ambil nilai tokennya
kita validasi tokenya
kita ambil user_id nya
ambil user dari db berdasarkan user_id lewat service
kita set context isinya user

*/
