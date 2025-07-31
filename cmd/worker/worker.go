package worker

import (
	"github.com/rms-diego/rinha-backend-2025/internal/database"
	"github.com/rms-diego/rinha-backend-2025/internal/service"
	"github.com/rms-diego/rinha-backend-2025/pkg/pubsub"
)

const WORKERS_POOL = 10

func Init() {
	ps := service.NewPaymentService(database.Db)
	pubsub.NewPubSub()

	for range WORKERS_POOL {
		go pubsub.Queue.Subscribe(ps.CreatePayment)
	}
}
