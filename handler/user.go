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

//tangkap input dari user
//map input dari user ke struct UserRegisterInput
//struct di atas kita passing sebagai parameter service
