package controller

import (
	"github.com/gin-gonic/gin"
	"pokedex/repository"
	"strconv"
)

type PokedexController struct {
	repo repository.PokedexRepository
}

func NewPokedexController(repo repository.PokedexRepository) PokedexController {
	return PokedexController{repo}
}

func (ctrl PokedexController) GetPokedex(ctx *gin.Context) {
	ctx.JSON(200, ctrl.repo.FindPokedexes())
	ctx.Writer.Flush()
}

func (ctrl PokedexController) GetPokedexByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, nil)
		ctx.Writer.Flush()
		return
	}
	ctx.JSON(200, ctrl.repo.FindPokedexByID(id))
	ctx.Writer.Flush()
}
