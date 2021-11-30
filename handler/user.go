package handler

import (
	"bwastartup/auth"
	"bwastartup/helper"
	"bwastartup/user"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
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
	token, err := h.authService.GenerateToken(newUser.ID)
	if err != nil {
		response := helper.ApiResponse("Register Account Failed", http.StatusBadRequest, "fail", err.Error())

		c.JSON(http.StatusBadRequest, response)
		return

	}
	formatResponse := user.FormatUser(newUser, token)

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
	token, err := h.authService.GenerateToken(userLogin.ID)
	if err != nil {
		// fmt.Println("[RESPONSE]", err.Error())

		response := helper.ApiResponse("Login Failed Token", http.StatusBadRequest, "fail", err.Error())

		c.JSON(http.StatusBadRequest, response)
		return

	}
	formatResponse := user.FormatUser(userLogin, token)

	response := helper.ApiResponse("Login Success", http.StatusOK, "success", formatResponse)
	c.JSON(http.StatusOK, response)

}

func (h *userHandler) IsEmailAvailable(c *gin.Context) {
	var email user.EmailUserInput
	err := c.ShouldBindJSON(&email)

	if err != nil {
		messageError := helper.FormatErrorValidation(err)
		response := helper.ApiResponse("Email Checking Failed", http.StatusBadRequest, "error", messageError)
		c.JSON(http.StatusOK, response)
		return
	}

	checkedEmail, err := h.userService.IsEmailAvailable(email)
	if err != nil {
		messageError := gin.H{"errors": "Server Error"}
		response := helper.ApiResponse("Email Checking Failed", http.StatusBadRequest, "error", messageError)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	messageError := gin.H{"is_avalible": checkedEmail}
	response := helper.ApiResponse("Email Checking Success", http.StatusOK, "success", messageError)
	c.JSON(http.StatusOK, response)

}

func (h *userHandler) UploadAvatar(c *gin.Context) {
	/*
		1.input dari user
		2.simpan gambarnya di folder "images/"
		3.di service kita panggil repo
		4.JWT
		5.repo ambil datauser
		6.repo update user data simpan lokasi file
	*/

	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.ApiResponse("Failed Upload Avatar", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// fmt.Println("[LOG]", *file)

	currenUser := c.MustGet("currentUser").(user.User)
	path := fmt.Sprintf("images/%d-%s", currenUser.ID, file.Filename)
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.ApiResponse("Failed Upload Avatar", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.userService.UploadAvatar(currenUser.ID, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.ApiResponse("Failed Upload Avatar", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.ApiResponse("Success Upload Avatar", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}
