package service

import (
	"encoding/json"

	"github.com/thalisonh/rinha-go/internal/domain/model"
)

type StreamSender interface {
	Send(payload []byte) error
}

type ExampleService struct {
	streamSender StreamSender
}

func NewExampleService(sender StreamSender) *ExampleService {
	return &ExampleService{streamSender: sender}
}

func (s *ExampleService) ExampleBusinessLogic() string {
	return "business logic result"
}

func (s *ExampleService) CreatePayment(input model.CreatePaymentInput) model.CreatePaymentOutput {
	payload, _ := json.Marshal(input)
	_ = s.streamSender.Send(payload)
	// Lógica de negócio fictícia
	return model.CreatePaymentOutput{
		ID:     "1",
		Status: "created",
		Amount: input.Amount,
	}
}

func (s *ExampleService) GetPaymentSummary() model.PaymentSummaryOutput {
	// Lógica de negócio fictícia
	return model.PaymentSummaryOutput{
		TotalPayments: 1,
		TotalAmount:   100.0,
	}
}
