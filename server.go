package main

import (
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"
	"log"
	"pokedex/controller"
	"pokedex/repository"
)

func api(db *gorm.DB, ch *cache.Cache, port string, favicon string) {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.StaticFile("/favicon.ico", favicon)
	router.Use(func (ctx *gin.Context) { ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*") })

	v1 := router.Group("/api/v1")
	repo := repository.NewPokemonRepository(db, ch)
	ctrl := controller.NewPokemonController(repo)
	v1.GET("/pokedex", ctrl.GetPokedex)
	v1.GET("/pokemon", ctrl.GetPokemon)
	v1.GET("/pokemon/:id", ctrl.GetPokemonByID)
	v1.GET("/pokemon/:id/variants", ctrl.GetPokemonVariantsByID)
	v1.GET("/pokemon/:id/variants/:vid", ctrl.GetPokemonVariantByID)

	log.Fatal(router.Run("0.0.0.0:" + port))
}