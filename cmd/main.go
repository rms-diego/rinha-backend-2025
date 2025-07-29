package main

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rms-diego/rinha-backend-2025/internal/config"
	"github.com/rms-diego/rinha-backend-2025/internal/database"
	routes "github.com/rms-diego/rinha-backend-2025/internal/route"
	"github.com/rms-diego/rinha-backend-2025/internal/service"
	"github.com/rms-diego/rinha-backend-2025/internal/shared"
)

func main() {
	if err := config.NewConfig(); err != nil {
		panic(err.Error())
	}

	ctx := context.Background()
	if err := database.Init(ctx); err != nil {
		panic(err.Error())
	}

	shared.NewPubSub()
	shared.Queue.Subscribe(service.CreatePayment)

	app := gin.New()
	routes.Init(app)
	app.Run(fmt.Sprintf(":%v", config.Env.PORT))
}
