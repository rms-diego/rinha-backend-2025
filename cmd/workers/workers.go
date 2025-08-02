package workers

import (
	"github.com/rms-diego/rinha-backend-2025/internal/database"
	"github.com/rms-diego/rinha-backend-2025/internal/service"
	"github.com/rms-diego/rinha-backend-2025/pkg/gateway"
	"github.com/rms-diego/rinha-backend-2025/pkg/pubsub"
)

const DEFAULT_WORKERS = 10
const FALLBACK_WORKERS = 5

func Init() {
	g := gateway.NewGateway(service.NewPaymentService(database.Db))

	pubsub.Queue.Subscribe(
		g.PaymentProcessor,
		g.PaymentProcessor,
		DEFAULT_WORKERS,
		FALLBACK_WORKERS,
	)
}
