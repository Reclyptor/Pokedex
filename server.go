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
	router.Use(func(ctx *gin.Context) { ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*") })

	pokemoncontroller := controller.NewPokemonController(repository.NewPokemonRepository(db, ch))
	pokedexcontroller := controller.NewPokedexController(repository.NewPokedexRepository(db, ch))

	v1 := router.Group("/api/v1")
	v1.GET("/pokedex", pokedexcontroller.GetPokedex)
	v1.GET("/pokedex/:id", pokedexcontroller.GetPokedexByID)
	v1.GET("/pokemon", pokemoncontroller.GetPokemon)
	v1.GET("/pokemon/:id", pokemoncontroller.GetPokemonByID)
	v1.GET("/pokemon/:id/variants", pokemoncontroller.GetPokemonVariantsByID)
	v1.GET("/pokemon/:id/variants/:vid", pokemoncontroller.GetPokemonVariantByID)

	log.Fatal(router.Run("0.0.0.0:" + port))
}
