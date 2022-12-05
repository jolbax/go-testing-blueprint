package greeting

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"go-testing-blueprint/internal/model"
	"net/http"
)

type HTTP struct {
	service GreetingService
}

func NewHTTP(gs GreetingService) *HTTP {
	return &HTTP{gs}
}

func (ctrl HTTP) Greet(ctx *gin.Context) {

	log.Info().Msg("Greet endpoint was called")
	ctx.JSON(http.StatusOK, ctrl.service.GetGreeting(ctx.Query("name")))
}

func (ctrl HTTP) GreetPost(ctx *gin.Context) {
	greeting := model.Greeting{}
	err := ctx.BindJSON(&greeting)

	if err != nil {
		log.Error().Msg("Error calling the GreetPost endpoint")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad formed request"})
		return
	}

	log.Info().Msg("GreetPost endpoint was called")
	ctx.JSON(http.StatusOK, ctrl.service.PostGreeting(greeting))

}
