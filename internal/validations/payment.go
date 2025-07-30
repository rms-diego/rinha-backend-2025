package validations

import "github.com/google/uuid"

type CreatePayment struct {
	Amount        float64   `json:"amount" binding:"required"`
	CorrelationId uuid.UUID `json:"correlationId" binding:"required"`
}

type Summary struct {
	TotalRequests int64   `json:"totalRequests" db:"TotalRequests"`
	TotalAmount   float64 `json:"totalAmount" db:"TotalAmount"`
}

type PaymentSummary struct {
	Default  Summary `json:"default"`
	Fallback Summary `json:"fallback"`
}
