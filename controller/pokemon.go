package controller

import (
	"github.com/gin-gonic/gin"
	"pokedex/repository"
	"strconv"
)

type PokemonController struct {
	repo repository.PokemonRepository
}

func NewPokemonController(repo repository.PokemonRepository) PokemonController {
	return PokemonController{repo}
}

func (ctrl PokemonController) GetPokedex(ctx *gin.Context) {
	ctx.JSON(200, ctrl.repo.FindPokedex())
	ctx.Writer.Flush()
}

func (ctrl PokemonController) GetPokemon(ctx *gin.Context) {
	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		limit = -1
	}

	offset, err := strconv.Atoi(ctx.Query("offset"))
	if err != nil {
		offset = -1
	}

	if limit < 0 || offset < 0 {
		ctx.JSON(200, ctrl.repo.FindAllPokemon())
	} else {
		ctx.JSON(200, ctrl.repo.FindPokemon(limit, offset))
	}

	ctx.Writer.Flush()
}

func (ctrl PokemonController) GetPokemonByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, nil)
		ctx.Writer.Flush()
		return
	}
	ctx.JSON(200, ctrl.repo.FindPokemonByID(id))
	ctx.Writer.Flush()
}

func (ctrl PokemonController) GetPokemonVariantsByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, nil)
		ctx.Writer.Flush()
		return
	}
	variants := ctrl.repo.FindPokemonVariantsByID(id)
	ctx.JSON(200, variants)
	ctx.Writer.Flush()
}

func (ctrl PokemonController) GetPokemonVariantByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	vid, err := strconv.Atoi(ctx.Param("vid"))
	if err != nil {
		ctx.JSON(400, nil)
		ctx.Writer.Flush()
		return
	}
	variant := ctrl.repo.FindPokemonVariantByID(id, vid)
	ctx.JSON(200, variant)
	ctx.Writer.Flush()
}
