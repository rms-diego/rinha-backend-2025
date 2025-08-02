package gateway

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/rms-diego/rinha-backend-2025/internal/config"
	"github.com/rms-diego/rinha-backend-2025/internal/service"
	"github.com/rms-diego/rinha-backend-2025/internal/validations"
	"github.com/rms-diego/rinha-backend-2025/pkg/pubsub"
)

type gatewayImplementation interface {
	PaymentProcessor(validations.Message, string) error
}

type gateway struct {
	paymentService service.PaymentServiceInterface
}

type sendPaymentGateway struct {
	Amount        float64 `json:"amount"`
	CorrelationId string  `json:"correlationId"`
	RequestedAt   string  `json:"requestedAt"`
}

func NewGateway(paymentService service.PaymentServiceInterface) gatewayImplementation {
	return &gateway{
		paymentService: paymentService,
	}
}

func (g *gateway) PaymentProcessor(msg validations.Message, processorType string) error {
	var p string

	switch processorType {
	case pubsub.DEFAULT_QUEUE:
		p = config.Env.PAYMENT_PROCESSOR_DEFAULT_URL

	default:
		p = config.Env.PAYMENT_PROCESSOR_FALLBACK_URL
	}

	fmt.Println("message receive on gateway: ", msg.CorrelationId, " - processor type: ", processorType)
	now := time.Now().Format(time.RFC3339)

	payload := sendPaymentGateway{
		Amount:        msg.Amount,
		CorrelationId: msg.CorrelationId,
		RequestedAt:   now,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%v/payments", p)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send payment: %s", resp.Status)
	}

	fmt.Println("payment processed successfully on processor: ", processorType)
	cp := validations.CreatePayment{
		Amount:        msg.Amount,
		CorrelationId: msg.CorrelationId,
		RequestedAt:   now,
	}

	go g.paymentService.CreatePayment(&cp, processorType)

	return nil
}
