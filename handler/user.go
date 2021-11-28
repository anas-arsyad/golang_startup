package handler

import (
	"bwastartup/helper"
	"bwastartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	//tangkap input dari user
	//map input dari user ke struct UserRegisterInput
	//struct di atas kita passing sebagai parameter service
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {

		messageError := helper.FormatErrorValidation(err)
		response := helper.ApiResponse("Register Account Failed", http.StatusBadRequest, "fail", messageError)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)

	if err != nil {
		response := helper.ApiResponse("Register Account Failed", http.StatusBadRequest, "fail", err.Error())

		c.JSON(http.StatusBadRequest, response)
		return

	}
	formatResponse := user.FormatUser(newUser, "token")

	response := helper.ApiResponse("Account has been Registered", http.StatusOK, "success", formatResponse)
	// fmt.Println("[RESPONSE]", response)
	c.JSON(http.StatusOK, response)

}

func (h *userHandler) LoginUser(c *gin.Context) {
	var input user.LoginUserInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		messageError := helper.FormatErrorValidation(err)
		response := helper.ApiResponse("Login Failed", http.StatusBadRequest, "failed", messageError)
		c.JSON(http.StatusOK, response)
		return
	}
	userLogin, err := h.userService.LoginUser(input)
	if err != nil {
		messageError := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("Login Account Failed", http.StatusBadRequest, "fail", messageError)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatResponse := user.FormatUser(userLogin, "token")

	response := helper.ApiResponse("Login Success", http.StatusOK, "success", formatResponse)
	// fmt.Println("[RESPONSE]", response)
	c.JSON(http.StatusOK, response)

}
