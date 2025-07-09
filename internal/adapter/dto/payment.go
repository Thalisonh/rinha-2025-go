package dto

type CreatePaymentRequest struct {
	Amount        float64 `json:"amount" binding:"required"`
	CorrelationID string  `json:"correlationId" binding:"required"`
}
