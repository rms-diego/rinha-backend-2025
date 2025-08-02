package validations

type CreatePayment struct {
	Amount        float64
	CorrelationId string
	RequestedAt   string
}

type Summary struct {
	TotalRequests int64   `json:"totalRequests" db:"TotalRequests"`
	TotalAmount   float64 `json:"totalAmount" db:"TotalAmount"`
}

type PaymentSummary struct {
	Default  Summary `json:"default"`
	Fallback Summary `json:"fallback"`
}
