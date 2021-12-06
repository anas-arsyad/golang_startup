package handler

import (
	"bwastartup/campaign"
	"bwastartup/helper"
	"bwastartup/user"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	campaignService campaign.Service
}

func NewCampaignHandler(c campaign.Service) *campaignHandler {
	return &campaignHandler{c}
}

func (h *campaignHandler) GetCampaign(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Query("user_id"))

	dataCamp, err := h.campaignService.GetCampaigns(userId)
	if err != nil {
		response := helper.ApiResponse("Failed Get Campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatJson := campaign.FormatCampaigns(dataCamp)
	response := helper.ApiResponse("Success Get Campaign", http.StatusOK, "success", formatJson)

	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) GetCampaignById(c *gin.Context) {
	input := campaign.GetCampaignById{}

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.ApiResponse("Failed Get Campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	campaignDetail, err := h.campaignService.GetCampaignById(input)
	if err != nil {
		response := helper.ApiResponse("Failed Get Campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Success Get Campaign Detail", http.StatusOK, "success", campaign.FormatCampaignById(campaignDetail))
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) CreateCampaign(c *gin.Context) {
	input := campaign.CreateCampaignInput{}

	err := c.ShouldBindJSON(&input)
	if err != nil {
		messageError := helper.FormatErrorValidation(err)
		response := helper.ApiResponse("Failed Create Campaign", http.StatusBadRequest, "error", messageError)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	currenUser := c.MustGet("currentUser").(user.User)
	input.User = currenUser
	newCampaign, err := h.campaignService.CreateCampaign(input)
	if err != nil {
		response := helper.ApiResponse("Failed Create Campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Success Creater Campaign", http.StatusOK, "success", campaign.FormatCampaign(newCampaign))
	c.JSON(http.StatusOK, response)
}
func (h *campaignHandler) UpdateCampaign(c *gin.Context) {
	input := campaign.CreateCampaignInput{}
	inputId := campaign.GetCampaignById{}

	err := c.ShouldBindUri(&inputId)
	if err != nil {
		response := helper.ApiResponse("Failed Get Campaign 1", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	err = c.ShouldBindJSON(&input)
	if err != nil {
		messageError := helper.FormatErrorValidation(err)
		response := helper.ApiResponse("Failed Update Campaign 2", http.StatusBadRequest, "error", messageError)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	currenUser := c.MustGet("currentUser").(user.User)
	input.User = currenUser
	updateCampaign, err := h.campaignService.UpdateCampaign(inputId, input)
	if err != nil {
		response := helper.ApiResponse("Failed Update Campaign 3", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Success Update Campaign", http.StatusOK, "success", campaign.FormatCampaign(updateCampaign))
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) UploadImage(c *gin.Context) {
	input := campaign.CreateCampaignImageInput{}

	err := c.ShouldBind(&input)
	if err != nil {
		messageError := helper.FormatErrorValidation(err)
		response := helper.ApiResponse("Failed Upload Campaign ", http.StatusBadRequest, "error", messageError)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.ApiResponse("Failed Upload Campaign Image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currenUser := c.MustGet("currentUser").(user.User)
	input.User = currenUser
	path := fmt.Sprintf("images/%d-%s", currenUser.ID, file.Filename)
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.ApiResponse("Failed Upload Campaign Image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.campaignService.SaveCampaignImage(input, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.ApiResponse("Failed Upload Campaign Image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.ApiResponse("Success Upload Campaign Image", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}
