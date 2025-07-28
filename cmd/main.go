package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rms-diego/rinha-backend-2025/internal/config"
)

func main() {
	if err := config.NewConfig(); err != nil {
		panic(err.Error())
	}

	app := gin.New()
	app.Run(fmt.Sprintf(":%v", config.Env.PORT))
}
