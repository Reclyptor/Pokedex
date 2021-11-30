package repository

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"pokedex/model"
)

func FindAllPokemon(db *gorm.DB) model.Pokedex {
	var pokedex model.Pokedex
	_ = db.Preload("Types").Preload("Physique").Preload("Abilities").Preload("Statistics").Preload("Training").Preload("Images").Find(&pokedex)
	return pokedex
}

func FindPokemonByID(db *gorm.DB, id int) model.Pokemon {
	var pokemon model.Pokemon
	_ = db.Where(model.Pokemon{PokedexID: id, IsDefault: true}).Preload(clause.Associations).First(&pokemon)
	return pokemon
}

func FindPokemonVariantsByID(db *gorm.DB, id int) []model.Pokemon {
	var variants []model.Pokemon
	_ = db.Where(model.Pokemon{PokedexID: id, IsDefault: false}, "PokedexID", "IsDefault").Preload("Types").Preload("Physique").Preload("Abilities").Preload("Statistics").Preload("Training").Preload("Images").Find(&variants)
	return variants
}

func FindPokemonVariantByID(db *gorm.DB, id int, vid int) *model.Pokemon {
	var variants []model.Pokemon
	_ = db.Where(model.Pokemon{PokedexID: id, IsDefault: false}, "PokedexID", "IsDefault").Preload(clause.Associations).Find(&variants)
	if len(variants) > vid - 1 {
		return &variants[vid - 1]
	} else {
		return nil
	}
}