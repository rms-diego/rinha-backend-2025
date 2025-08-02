package validations

type Message struct {
	Amount        float64
	CorrelationId string
}

func NewMessage(payment CreatePayment) Message {
	return Message{
		Amount:        payment.Amount,
		CorrelationId: payment.CorrelationId,
	}
}
