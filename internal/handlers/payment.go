package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/rms-diego/rinha-backend-2025/internal/service"
	"github.com/rms-diego/rinha-backend-2025/internal/validations"
	"github.com/rms-diego/rinha-backend-2025/pkg/pubsub"
)

type paymentHandlerInterface interface {
	CreatePayment(c *gin.Context)
	PaymentsSummary(c *gin.Context)
}

type paymentHandler struct {
	service service.PaymentServiceInterface
}

func NewPaymentHandler(paymentService service.PaymentServiceInterface) paymentHandlerInterface {
	return &paymentHandler{
		service: paymentService,
	}
}

func (h *paymentHandler) CreatePayment(c *gin.Context) {
	body := &validations.CreatePayment{}

	if err := c.ShouldBindJSON(body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	go pubsub.Queue.Publish(*body)

	c.JSON(204, nil)
}

func (h *paymentHandler) PaymentsSummary(c *gin.Context) {
	f := c.Query("from")
	t := c.Query("to")

	if f == "" || t == "" {
		c.JSON(400, gin.H{"error": "from and to query parameters are required"})
		return
	}

	result, err := h.service.ListPaymentsSummary(f, t)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, result)
}
