package fruits

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type HTTP struct {
	service FruitsService
}

func NewHTTP(fs FruitsService) *HTTP {
	return &HTTP{fs}
}

func (ctrl HTTP) Get(ctx *gin.Context) {

	log.Info().Msg("GetFruit endpoint was called")
	fr, err := ctrl.service.GetFruits()
	if err != nil {
		ctx.JSON(fr.Code, err.Error())
		return
	}
	ctx.JSON(fr.Code, fr.Fruits)
}

func (ctrl HTTP) Post(ctx *gin.Context) {
	log.Info().Msg("PostFruit endpoint was called")
	fr, err := ctrl.service.PostFruits()
	if err != nil {
		ctx.JSON(fr.Code, err.Error())
		return
	}
	ctx.JSON(fr.Code, fr.Fruits)
}
