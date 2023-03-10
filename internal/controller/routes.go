package controller

import (
	"github.com/gin-gonic/gin"
	"go-testing-blueprint/internal/controller/greeting"
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

}
