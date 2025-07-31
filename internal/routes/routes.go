package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rms-diego/rinha-backend-2025/internal/database"
	"github.com/rms-diego/rinha-backend-2025/internal/handlers"
	"github.com/rms-diego/rinha-backend-2025/internal/service"
)

func Init(app *gin.Engine) {
	ps := service.NewPaymentService(database.Db)
	dp := handlers.NewPaymentHandler(ps)

	INSTANCE_ID := uuid.New()
	app.GET("/payments/service-health", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "server is running", "instanceId": INSTANCE_ID})
	})

	app.POST("/payments", dp.CreatePayment)
	app.GET("/payments-summary", dp.PaymentsSummary)
}
