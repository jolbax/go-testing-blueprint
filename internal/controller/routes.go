package controller

import (
	"github.com/gin-gonic/gin"
	"go-testing-blueprint/internal/controller/fruits"
	"go-testing-blueprint/internal/controller/greeting"
	fruits_client "go-testing-blueprint/pkg/fruits-client"
)

type HTTP struct {
}

func NewHTTP() *HTTP {
	return &HTTP{}
}

func (ctrl HTTP) RegisterRoutes(ginEngine *gin.Engine) {
	greetService := greeting.NewGreetingService()
	greetCtrl := greeting.NewHTTP(greetService)

	api := ginEngine.Group("/api")

	api.GET("/greeting", greetCtrl.Greet)
	api.POST("/greetingPost", greetCtrl.GreetPost)

	fruitsClient := fruits_client.NewFruitsClient()
	fruitsService := fruits.NewFruitsService(fruitsClient)
	fruitsCtrl := fruits.NewHTTP(fruitsService)

	api.GET("/getFruits", fruitsCtrl.Get)
	api.POST("/postFruits", fruitsCtrl.Post)
}
