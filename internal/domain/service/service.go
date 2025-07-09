package service

import "github.com/thalisonh/rinha-go/internal/domain/model"

// Aqui ficam as regras de negócio (domain service)

type ExampleService struct{}

func NewExampleService() *ExampleService {
	return &ExampleService{}
}

func (s *ExampleService) ExampleBusinessLogic() string {
	return "business logic result"
}

func (s *ExampleService) CreatePayment(input model.CreatePaymentInput) model.CreatePaymentOutput {
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
