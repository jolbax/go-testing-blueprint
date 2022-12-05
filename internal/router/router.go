package router

import "github.com/gin-gonic/gin"

type Registerable interface {
	RegisterRoutes(engine *gin.Engine)
}

func NewRouter(registerables ...Registerable) *gin.Engine {
	ginEngine := gin.New()

	for _, registerable := range registerables {
		registerable.RegisterRoutes(ginEngine)
	}
	return ginEngine
}
