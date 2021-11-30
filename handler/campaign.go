package handler

import (
	"bwastartup/campaign"
	"bwastartup/helper"
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
