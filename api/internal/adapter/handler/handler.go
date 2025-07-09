package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thalisonh/rinha-go/internal/adapter/dto"
	"github.com/thalisonh/rinha-go/internal/domain/model"
	"github.com/thalisonh/rinha-go/internal/domain/service"
)

type Handler struct {
	Service *service.ExampleService
}

func RegisterRoutes(r *gin.Engine, h *Handler) {
	r.GET("/ping", h.Ping)
	r.GET("/business", h.Business)
	r.POST("/payments", h.CreatePayment)
	r.GET("/payments-summary", h.PaymentSummary)
}

func (h *Handler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func (h *Handler) Business(c *gin.Context) {
	result := h.Service.ExampleBusinessLogic()
	c.JSON(http.StatusOK, gin.H{"result": result})
}

func (h *Handler) CreatePayment(c *gin.Context) {
	var req dto.CreatePaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input := model.CreatePaymentInput{
		Amount:        req.Amount,
		CorrelationID: req.CorrelationID,
	}

	output := h.Service.CreatePayment(input)
	c.JSON(http.StatusCreated, output)
}

func (h *Handler) PaymentSummary(c *gin.Context) {
	output := h.Service.GetPaymentSummary()
	c.JSON(http.StatusOK, output)
}
