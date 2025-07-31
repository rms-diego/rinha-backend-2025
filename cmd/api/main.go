package main

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rms-diego/rinha-backend-2025/cmd/worker"
	"github.com/rms-diego/rinha-backend-2025/internal/config"
	"github.com/rms-diego/rinha-backend-2025/internal/database"
	routes "github.com/rms-diego/rinha-backend-2025/internal/routes"
)

func main() {
	if err := config.NewConfig(); err != nil {
		panic(err.Error())
	}

	ctx := context.Background()
	if err := database.Init(ctx); err != nil {
		panic(err.Error())
	}

	worker.Init()

	app := gin.New()
	routes.Init(app)
	fmt.Println("Server is running 🚀")
	app.Run(fmt.Sprintf(":%v", config.Env.PORT))
}
