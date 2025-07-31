package service

import (
	"fmt"
	"time"

	"github.com/doug-martin/goqu"
	"github.com/rms-diego/rinha-backend-2025/internal/database"
	"github.com/rms-diego/rinha-backend-2025/internal/validations"
)

type PaymentServiceInterface interface {
	CreatePayment(message validations.CreatePayment) error
	ListPaymentsSummary(from string, to string) (*validations.PaymentSummary, error)
}

type paymentService struct {
	database *goqu.Database
}

func NewPaymentService(database *goqu.Database) PaymentServiceInterface {
	return &paymentService{
		database: database,
	}
}

func (s *paymentService) CreatePayment(message validations.CreatePayment) error {
	_, err := database.Db.From("payments").Insert(goqu.Record{
		"amount":         message.Amount,
		"correlation_id": message.CorrelationId,
		"requested_at":   time.Now(),
	}).Exec()

	if err != nil {
		return err
	}

	return nil
}

func (s *paymentService) ListPaymentsSummary(from string, to string) (*validations.PaymentSummary, error) {
	var summary validations.Summary

	_, err := database.Db.
		From("payments").
		Select(
			goqu.COUNT("requested_at").As("TotalRequests"),
			goqu.SUM("amount").As("TotalAmount"),
		).
		Where(
			goqu.I("requested_at").Gte(from),
			goqu.I("requested_at").Lte(to),
		).
		ScanStruct(&summary)

	if err != nil {
		return nil, fmt.Errorf("erro ao executar query: %w", err)
	}

	return &validations.PaymentSummary{
		Default: summary,
		Fallback: validations.Summary{
			TotalRequests: 0,
			TotalAmount:   0,
		},
	}, nil
}
