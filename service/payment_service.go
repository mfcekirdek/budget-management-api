package service

import (
	"context"
	"gitlab.com/mfcekirdek/budget-management-api/model"
	"gitlab.com/mfcekirdek/budget-management-api/repository"
	"go.uber.org/zap"
)

type PaymentService interface {
	SavePayment(ctx context.Context, payment *model.Payment) (*model.Payment, error)
	GetAllPayments(ctx context.Context) ([]model.Payment, error)
}

type paymentService struct {
	logger            *zap.Logger
	paymentRepository repository.PaymentRepository
}

func NewPaymentService(logger *zap.Logger, repository repository.PaymentRepository) PaymentService {
	return &paymentService{logger: logger, paymentRepository: repository}
}

func (p paymentService) SavePayment(ctx context.Context, payment *model.Payment) (*model.Payment, error) {
	_, err := p.paymentRepository.SaveNewPayment(ctx, payment)
	if err != nil {
		return nil, err
	}
	return payment, nil
}

func (p paymentService) GetAllPayments(ctx context.Context) ([]model.Payment, error) {
	payments, err := p.paymentRepository.GetAllPayments(ctx)
	if err != nil {
		return nil, err
	}
	return payments, nil
}
