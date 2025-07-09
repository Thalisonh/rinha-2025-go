package service

import (
	"encoding/json"

	"github.com/thalisonh/rinha-go/internal/domain/model"
)

type StreamSender interface {
	Send(payload []byte) error
	Get() (string, error)
	Delete(id string) error
}

type ExampleService struct {
	streamSender StreamSender
}

func NewExampleService(sender StreamSender) *ExampleService {
	return &ExampleService{streamSender: sender}
}

func (s *ExampleService) CreatePayment(input model.CreatePaymentInput) *model.CreatePaymentOutput {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil
	}

	err = s.streamSender.Send(payload)
	if err != nil {
		return nil
	}
	// Lógica de negócio fictícia
	return &model.CreatePaymentOutput{
		ID:     "1",
		Status: "created",
		Amount: input.Amount,
	}
}

func (s *ExampleService) GetPaymentSummary() model.PaymentSummaryOutput {
	summary, err := s.streamSender.Get()
	if err != nil {
		return model.PaymentSummaryOutput{}
	}

	t := model.PaymentSummaryOutput{}
	json.Unmarshal([]byte(summary), &t)

	return model.PaymentSummaryOutput{
		TotalPayments: t.TotalPayments,
		TotalAmount:   t.TotalAmount,
	}
}
