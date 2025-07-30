package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rms-diego/rinha-backend-2025/internal/handlers"
)

func Init(app *gin.Engine) {
	INSTANCE_ID := uuid.New()

	app.GET("/payments/service-health", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "server is running", "instanceId": INSTANCE_ID})
	})

	app.POST("/payments", handlers.CreatePayment)
	app.GET("/payments-summary", handlers.PaymentsSummary)
}
