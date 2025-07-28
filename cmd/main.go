package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rms-diego/rinha-backend-2025/internal/config"
	routes "github.com/rms-diego/rinha-backend-2025/internal/route"
)

func main() {
	if err := config.NewConfig(); err != nil {
		panic(err.Error())
	}

	app := gin.New()
	routes.Init(app)
	app.Run(fmt.Sprintf(":%v", config.Env.PORT))
}
