package dto

type CreatePaymentRequest struct {
	Amount      float64 `json:"amount" binding:"required"`
	Description string  `json:"description" binding:"required"`
}
