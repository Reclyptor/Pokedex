package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"pokedex/controller"
)

func api(db *gorm.DB, port string, favicon string) {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.StaticFile("/favicon.ico", favicon)

	router.Use(func (ctx *gin.Context) { ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*") })

	v1 := router.Group("/api/v1")
	v1.GET("/pokemon", controller.GetPokedex(db))
	v1.GET("/pokemon/:id", controller.GetPokemonByID(db))
	v1.GET("/pokemon/:id/variants", controller.GetPokemonVariantsByID(db))
	v1.GET("/pokemon/:id/variants/:vid", controller.GetPokemonVariantByID(db))

	log.Fatal(router.Run("0.0.0.0:" + port))
}