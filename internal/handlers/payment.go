package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/rms-diego/rinha-backend-2025/internal/shared"
	"github.com/rms-diego/rinha-backend-2025/internal/validations"
)

func CreatePaymentHandler(c *gin.Context) {
	body := &validations.CreatePaymentRequest{}

	if err := c.ShouldBindJSON(body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	go func() {
		shared.Queue.Publish(*body)
	}()
	c.JSON(204, nil)
}
