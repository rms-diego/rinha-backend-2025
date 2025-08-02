package pubsub

import (
	"fmt"

	"github.com/rms-diego/rinha-backend-2025/internal/validations"
)

const (
	DEFAULT_QUEUE  = "default"
	FALLBACK_QUEUE = "fallback"
)
const MAX_MESSAGES_PER_QUEUE = 20

type PubSub struct {
	defaultMessages  chan validations.Message
	fallbackMessages chan validations.Message
}

var Queue *PubSub

func NewPubSub() {
	Queue = &PubSub{
		defaultMessages:  make(chan validations.Message, MAX_MESSAGES_PER_QUEUE),
		fallbackMessages: make(chan validations.Message, MAX_MESSAGES_PER_QUEUE),
	}
}

func (p *PubSub) Publish(message validations.Message, queueType string) {
	switch queueType {
	case DEFAULT_QUEUE:
		Queue.defaultMessages <- message
	case FALLBACK_QUEUE:
		Queue.fallbackMessages <- message
	}
}

func (p *PubSub) Subscribe(
	defaultProcessor func(message validations.Message, processorType string) error,
	fallbackProcessor func(message validations.Message, processorType string) error,
	defaultWorkers int,
	fallbackWorkers int,
) {
	// Workers default queue
	for range defaultWorkers {
		go func() {
			for msg := range p.defaultMessages {
				if err := defaultProcessor(msg, DEFAULT_QUEUE); err != nil {
					p.Publish(msg, FALLBACK_QUEUE)
				}
			}
		}()
	}

	// Workers fallback queue
	for range fallbackWorkers {
		go func() {
			for msg := range p.fallbackMessages {
				if err := fallbackProcessor(msg, FALLBACK_QUEUE); err != nil {
					fmt.Println("Error processing message in fallback queue:", err)

					continue
				}
			}
		}()
	}
}
