package service

import (
	"fmt"
	"time"

	"github.com/doug-martin/goqu"
	"github.com/rms-diego/rinha-backend-2025/internal/database"
	"github.com/rms-diego/rinha-backend-2025/internal/validations"
)

func CreatePayment(message validations.CreatePaymentRequest) error {
	fmt.Println("Processing payment:", message)

	_, err := database.Db.From("payments").Insert(goqu.Record{
		"amount":        message.Amount,
		"correlationId": message.CorrelationId,
		"requestedAt":   time.Now(),
	}).Exec()

	if err != nil {
		return err
	}

	return nil
}
