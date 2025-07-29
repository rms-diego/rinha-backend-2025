package shared

import (
	"github.com/rms-diego/rinha-backend-2025/internal/validations"
)

type PubSub interface {
	Publish(topic string, message validations.CreatePaymentRequest)
	Subscribe(handler func(message validations.CreatePaymentRequest))
}

type pubSub struct {
	messages chan validations.CreatePaymentRequest
}

var Queue pubSub

func NewPubSub() {
	Queue = pubSub{
		messages: make(chan validations.CreatePaymentRequest, 1000_000),
	}
}

func (p *pubSub) Publish(message validations.CreatePaymentRequest) {
	Queue.messages <- message
}

func (p *pubSub) Subscribe(handler func(message validations.CreatePaymentRequest) error) error {
	go func() {
		for msg := range Queue.messages {
			if err := handler(msg); err != nil {
				continue
			}
		}
	}()

	return nil
}
