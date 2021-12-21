package handler

import (
	"bwastartup/helper"
	"bwastartup/transaction"
	"bwastartup/user"
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
	currenUser := c.MustGet("currentUser").(user.User)
	input.User = currenUser
	transactionRes, err := h.transactionService.GetTransactionCampaignById(input)
	if err != nil {
		response := helper.ApiResponse("Failed Get Campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	trsansactions := transaction.FormatCampaignTransactions(transactionRes)
	response := helper.ApiResponse("Success Get Campaign", http.StatusOK, "success", trsansactions)
	c.JSON(http.StatusOK, response)
}

func (h *transactionHandler) GetUserTransactions(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)

	userId := currentUser.ID
	transactions, err := h.transactionService.GetTransactionUserById(userId)
	if err != nil {
		response := helper.ApiResponse("Failed Get User Transaction Campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := transaction.FormatUserTransactions(transactions)
	response := helper.ApiResponse("Failed Get User Transaction Campaign", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)

}

func (h *transactionHandler) CreateTransaction(c *gin.Context) {
	input := transaction.CreateTransactionInput{}
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatErrorValidation(err)
		errMessage := gin.H{"error": errors}
		response := helper.ApiResponse("Failed Save Transaction", http.StatusBadRequest, "error", errMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	currenUser := c.MustGet("currentUser").(user.User)
	input.User = currenUser
	newTransaction, err := h.transactionService.CreateTransaction(input)
	if err != nil {
		response := helper.ApiResponse("Failed Save Transaction", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := transaction.FormatTransaction(newTransaction)
	response := helper.ApiResponse("Success Save Transaction", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)

}
