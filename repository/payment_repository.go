package repository

import (
	"context"
	"fmt"
	"github.com/couchbase/gocb/v2"
	"github.com/google/uuid"
	"gitlab.com/mfcekirdek/budget-management-api/model"
	"go.uber.org/zap"
)

type PaymentRepository interface {
	SaveNewPayment(ctx context.Context, payment *model.Payment) (string, error)
	GetPaymentFromDocumentId(documentId string) (*model.Payment, error)
	GetAllPayments(ctx context.Context) ([]model.Payment, error)
}

type paymentRepository struct {
	logger     *zap.Logger
	collection *gocb.Collection
	cluster    *gocb.Cluster
}

func NewPaymentRepository(logger *zap.Logger, collection *gocb.Collection, cluster *gocb.Cluster) PaymentRepository {
	return &paymentRepository{logger: logger, collection: collection, cluster: cluster}
}

func (r paymentRepository) GetPaymentFromDocumentId(documentId string) (*model.Payment, error) {
	result := model.Payment{}
	getResult, err := r.collection.Get(documentId, &gocb.GetOptions{})
	if err != nil {
		return nil, err
	}
	if getResult != nil {
		if err := getResult.Content(&result); err != nil {
			panic(err)
		}
	}
	return &result, nil
}

func (r paymentRepository) SaveNewPayment(ctx context.Context, payment *model.Payment) (string, error) {
	documentId := uuid.New().String()
	_, err := r.collection.Insert(documentId, &payment, nil)
	if err != nil {
		return "", err
	}
	return documentId, nil
}

func (p paymentRepository) GetAllPayments(ctx context.Context) ([]model.Payment, error) {
	result := make([]model.Payment, 0)
	query := fmt.Sprintf("SELECT %s.* FROM %s", p.collection.Bucket().Name(), p.collection.Bucket().Name())
	rows, err := p.cluster.Query(query, nil)
	if err != nil {
		return nil, err
	}

	payment := model.Payment{}
	for rows.Next() {
		err := rows.Row(&payment)
		if err != nil {
			return nil, err
		}
		result = append(result, payment)
		payment = model.Payment{}
	}
	return result, nil
}
