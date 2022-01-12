package repository

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"pokedex/model"
)

type PokemonRepository struct {
	db    *gorm.DB
	cache *cache.Cache
}

func NewPokemonRepository(db *gorm.DB, cache *cache.Cache) PokemonRepository {
	return PokemonRepository{db, cache}
}

func (pokemonRepository PokemonRepository) FindAllPokemon() []model.Pokemon {
	if pokemon, found := pokemonRepository.cache.Get("pokemon"); found {
		return pokemon.([]model.Pokemon)
	} else {
		var pokemon []model.Pokemon
		_ = pokemonRepository.db.Where(model.Pokemon{IsDefault: true}).Preload("Types").Preload("Physique").Preload("Abilities").Preload("Statistics").Preload("Training").Preload("Images").Find(&pokemon)
		pokemonRepository.cache.Set("pokemon", pokemon, cache.DefaultExpiration)
		return pokemon
	}
}

func (pokemonRepository PokemonRepository) FindPokemon(limit int, offset int) []model.Pokemon {
	if pokemon, found := pokemonRepository.cache.Get(fmt.Sprintf("pokemon|%d|%d", limit, offset)); found {
		return pokemon.([]model.Pokemon)
	} else {
		var pokemon []model.Pokemon
		_ = pokemonRepository.db.Where(model.Pokemon{IsDefault: true}).Preload("Types").Preload("Physique").Preload("Abilities").Preload("Statistics").Preload("Training").Preload("Images").Limit(limit).Offset(offset).Find(&pokemon)
		pokemonRepository.cache.Set(fmt.Sprintf("pokemon|%d|%d", limit, offset), pokemon, cache.DefaultExpiration)
		return pokemon
	}
}

func (pokemonRepository PokemonRepository) FindPokemonByID(id int) model.Pokemon {
	var pokemon model.Pokemon
	_ = pokemonRepository.db.Where(model.Pokemon{PokedexNumber: id, IsDefault: true}).Preload(clause.Associations).Take(&pokemon)
	return pokemon
}

func (pokemonRepository PokemonRepository) FindPokemonVariantsByID(id int) []model.Pokemon {
	var variants []model.Pokemon
	_ = pokemonRepository.db.Where(model.Pokemon{PokedexNumber: id, IsDefault: false}, "PokedexID", "IsDefault").Preload("Types").Preload("Physique").Preload("Abilities").Preload("Statistics").Preload("Training").Preload("Images").Find(&variants)
	return variants
}

func (pokemonRepository PokemonRepository) FindPokemonVariantByID(id int, vid int) *model.Pokemon {
	var variants []model.Pokemon
	_ = pokemonRepository.db.Where(model.Pokemon{PokedexNumber: id, IsDefault: false}, "PokedexID", "IsDefault").Preload(clause.Associations).Find(&variants)
	if len(variants) > vid-1 {
		return &variants[vid-1]
	} else {
		return nil
	}
}
