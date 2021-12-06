package handler

import (
	"bwastartup/helper"
	"bwastartup/transaction"
	"net/http"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	transactionService transaction.Service
}

func NewTransactionHandler(transactionService transaction.Service) *transactionHandler {
	return &transactionHandler{transactionService}
}

func (h *transactionHandler) GetCampaignById(c *gin.Context) {
	// campaignId,_:=strconv.Atoi()
	input := transaction.GetCampaignIdInput{}
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.ApiResponse("Failed Get Campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	transaction, err := h.transactionService.GetTransactionCampaignById(input)
	if err != nil {
		response := helper.ApiResponse("Failed Get Campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Success Get Campaign", http.StatusOK, "success", transaction)
	c.JSON(http.StatusOK, response)
}
