package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"pokedex/repository"
	"strconv"
)

func GetPokedex(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pokedex := repository.FindAllPokemon(db)
		ctx.JSON(200, pokedex)
		ctx.Writer.Flush()
	}
}

func GetPokemonByID(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 0)
		if err != nil {
			ctx.JSON(400, nil)
			ctx.Writer.Flush()
			return
		}
		pokemon := repository.FindPokemonByID(db, int(id))
		ctx.JSON(200, pokemon)
		ctx.Writer.Flush()
	}
}

func GetPokemonVariantsByID(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 0)
		if err != nil {
			ctx.JSON(400, nil)
			ctx.Writer.Flush()
			return
		}
		variants := repository.FindPokemonVariantsByID(db, int(id))
		ctx.JSON(200, variants)
		ctx.Writer.Flush()
	}
}

func GetPokemonVariantByID(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 0)
		vid, err := strconv.ParseInt(ctx.Param("vid"), 10, 0)
		if err != nil {
			ctx.JSON(400, nil)
			ctx.Writer.Flush()
			return
		}
		variant := repository.FindPokemonVariantByID(db, int(id), int(vid))
		ctx.JSON(200, variant)
		ctx.Writer.Flush()
	}
}