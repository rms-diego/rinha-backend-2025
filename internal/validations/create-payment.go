package validations

import "github.com/google/uuid"

type CreatePaymentRequest struct {
	Amount        float64   `json:"amount" binding:"required"`
	CorrelationId uuid.UUID `json:"correlationId" binding:"required"`
}
