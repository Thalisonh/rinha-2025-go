package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thalisonh/rinha-go/internal/adapter/dto"
	"github.com/thalisonh/rinha-go/internal/domain/model"
	"github.com/thalisonh/rinha-go/internal/domain/service"
)

// Models de input/output

type CreatePaymentInput struct {
	Amount      float64 `json:"amount" binding:"required"`
	Description string  `json:"description" binding:"required"`
}

type CreatePaymentOutput struct {
	ID     string  `json:"id"`
	Status string  `json:"status"`
	Amount float64 `json:"amount"`
}

type PaymentSummaryOutput struct {
	TotalPayments int     `json:"total_payments"`
	TotalAmount   float64 `json:"total_amount"`
}

func RegisterRoutes(r *gin.Engine) {
	r.GET("/ping", Ping)
	r.GET("/business", Business)
	r.POST("/payments", CreatePayment)
	r.GET("/payments-summary", PaymentSummary)
}

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func Business(c *gin.Context) {
	svc := service.NewExampleService()
	result := svc.ExampleBusinessLogic()
	c.JSON(http.StatusOK, gin.H{"result": result})
}

func CreatePayment(c *gin.Context) {
	var req dto.CreatePaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input := model.CreatePaymentInput{
		Amount:      req.Amount,
		Description: req.Description,
	}

	svc := service.NewExampleService()
	output := svc.CreatePayment(input)
	c.JSON(http.StatusCreated, output)
}

func PaymentSummary(c *gin.Context) {
	svc := service.NewExampleService()
	output := svc.GetPaymentSummary()
	c.JSON(http.StatusOK, output)
}
