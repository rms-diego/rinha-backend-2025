package service

import (
	"fmt"
	"strings"
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
	sql := database.Db.From("payments").Insert(goqu.Ex{
		"amount":        message.Amount,
		"correlationId": message.CorrelationId,
		"requested_at":  time.Now(),
	}).Sql

	// Remove double quotes from the SQL string to avoid syntax errors
	cleanSQL := sanitizeSQL(sql)

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

func sanitizeSQL(sql string) string {
	return strings.ReplaceAll(sql, `"`, "")
}
