package pubsub

import (
	"fmt"

	"github.com/rms-diego/rinha-backend-2025/internal/validations"
)

type PubSub interface {
	Publish(message validations.CreatePayment)
	Subscribe(handler func(message validations.CreatePayment))
}

type pubSub struct {
	messages chan validations.CreatePayment
}

var Queue *pubSub

const MAX_MESSAGES = 100_000

func NewPubSub() {
	Queue = &pubSub{
		messages: make(chan validations.CreatePayment, MAX_MESSAGES),
	}
}

func (p *pubSub) Publish(message validations.CreatePayment) {
	Queue.messages <- message
}

func (p *pubSub) Subscribe(handler func(message validations.CreatePayment) error) {
	for msg := range Queue.messages {
		fmt.Println("Processing message: ", msg.CorrelationId)

		if err := handler(msg); err != nil {
			Queue.Publish(msg) // Re-publish the message if handler fails
			continue
		}
	}
}
