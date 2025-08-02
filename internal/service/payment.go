package service

import (
	"fmt"
	"strings"

	"github.com/doug-martin/goqu"
	"github.com/rms-diego/rinha-backend-2025/internal/database"
	"github.com/rms-diego/rinha-backend-2025/internal/validations"
	"github.com/rms-diego/rinha-backend-2025/pkg/pubsub"
)

type PaymentServiceInterface interface {
	CreatePayment(message *validations.CreatePayment, processedBy string) error
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

func (s *paymentService) CreatePayment(message *validations.CreatePayment, processedBy string) error {
	isDefaultProcessor := true

	switch processedBy {
	case pubsub.FALLBACK_QUEUE:
		isDefaultProcessor = false
	default:
		isDefaultProcessor = true
	}

	sql := database.Db.From("payments").Insert(goqu.Ex{
		"amount":               message.Amount,
		"correlationId":        message.CorrelationId,
		"requested_at":         message.RequestedAt,
		"is_default_processor": isDefaultProcessor,
	}).Sql

	cleanSQL := strings.ReplaceAll(sql, `"`, "")
	if _, err := database.Db.Exec(cleanSQL); err != nil {
		return err
	}

	return nil
}

func (s *paymentService) ListPaymentsSummary(from string, to string) (*validations.PaymentSummary, error) {
	var summary validations.Summary

	_, err := database.Db.
		From("payments").
		Select(
			goqu.COUNT(goqu.I("correlationid")).As("TotalRequests"),
			goqu.COALESCE(goqu.SUM(goqu.I("amount")), 0).As("TotalAmount"),
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
