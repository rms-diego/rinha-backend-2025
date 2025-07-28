package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Init(app *gin.Engine) {
	INSTANCE_ID := uuid.New()

	app.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "server is running", "instanceId": INSTANCE_ID})
	})
}
