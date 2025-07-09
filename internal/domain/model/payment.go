package model

type CreatePaymentInput struct {
	Amount      float64
	Description string
}

type CreatePaymentOutput struct {
	ID     string
	Status string
	Amount float64
}

type PaymentSummaryOutput struct {
	TotalPayments int
	TotalAmount   float64
}
