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
	type aggregationResult struct {
		IsDefaultProcessor bool    `db:"is_default_processor"`
		TotalRequests      int64   `db:"TotalRequests"`
		TotalAmount        float64 `db:"TotalAmount"`
	}

	var results []aggregationResult

	err := database.Db.
		From("payments").
		Select(
			goqu.I("is_default_processor"),
			goqu.COUNT(goqu.I("correlationid")).As("TotalRequests"),
			goqu.COALESCE(goqu.SUM(goqu.I("amount")), 0).As("TotalAmount"),
		).
		Where(
			goqu.I("requested_at").Gte(from),
			goqu.I("requested_at").Lte(to),
		).
		GroupBy(goqu.I("is_default_processor")).
		ScanStructs(&results)

	if err != nil {
		return nil, fmt.Errorf("erro ao executar query: %w", err)
	}

	var summary validations.PaymentSummary

	for _, data := range results {
		if data.IsDefaultProcessor {
			summary.Default = validations.Summary{
				TotalRequests: data.TotalRequests,
				TotalAmount:   data.TotalAmount,
			}
			continue
		}

		summary.Fallback = validations.Summary{
			TotalRequests: data.TotalRequests,
			TotalAmount:   data.TotalAmount,
		}
	}

	return &summary, nil

}
