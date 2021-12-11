package repository

import (
	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"pokedex/model"
)

type PokemonRepository struct {
	db *gorm.DB
	cache *cache.Cache
}

func NewPokemonRepository(db *gorm.DB, cache *cache.Cache) PokemonRepository {
	return PokemonRepository{db, cache}
}

func (pokemonRepository PokemonRepository) FindPokedex() model.Pokedex {
	if pokedex, found := pokemonRepository.cache.Get("pokedex"); found {
		return pokedex.(model.Pokedex)
	} else {
		var pokedex model.Pokedex
		_ = pokemonRepository.db.Preload("Types").Preload("Physique").Preload("Abilities").Preload("Statistics").Preload("Training").Preload("Images").Find(&pokedex)
		pokemonRepository.cache.Set("pokedex", pokedex, cache.DefaultExpiration)
		return pokedex
	}
}

func (pokemonRepository PokemonRepository) FindPokemon() []model.Pokemon {
	if pokemon, found := pokemonRepository.cache.Get("pokemon"); found {
		return pokemon.([]model.Pokemon)
	} else {
		var pokemon []model.Pokemon
		_ = pokemonRepository.db.Where(model.Pokemon{IsDefault: true}).Preload("Types").Preload("Physique").Preload("Abilities").Preload("Statistics").Preload("Training").Preload("Images").Find(&pokemon)
		pokemonRepository.cache.Set("pokemon", pokemon, cache.DefaultExpiration)
		return pokemon
	}
}

func (pokemonRepository PokemonRepository) FindPokemonByID(id int) model.Pokemon {
	var pokemon model.Pokemon
	_ = pokemonRepository.db.Where(model.Pokemon{PokedexID: id, IsDefault: true}).Preload(clause.Associations).Take(&pokemon)
	return pokemon
}

func (pokemonRepository PokemonRepository) FindPokemonVariantsByID(id int) []model.Pokemon {
	var variants []model.Pokemon
	_ = pokemonRepository.db.Where(model.Pokemon{PokedexID: id, IsDefault: false}, "PokedexID", "IsDefault").Preload("Types").Preload("Physique").Preload("Abilities").Preload("Statistics").Preload("Training").Preload("Images").Find(&variants)
	return variants
}

func (pokemonRepository PokemonRepository) FindPokemonVariantByID(id int, vid int) *model.Pokemon {
	var variants []model.Pokemon
	_ = pokemonRepository.db.Where(model.Pokemon{PokedexID: id, IsDefault: false}, "PokedexID", "IsDefault").Preload(clause.Associations).Find(&variants)
	if len(variants) > vid - 1 {
		return &variants[vid - 1]
	} else {
		return nil
	}
}