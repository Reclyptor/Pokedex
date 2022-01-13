package repository

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"pokedex/model"
)

type PokedexRepository struct {
	db    *gorm.DB
	cache *cache.Cache
}

func NewPokedexRepository(db *gorm.DB, cache *cache.Cache) PokedexRepository {
	return PokedexRepository{db, cache}
}

func (pokedexRepository PokedexRepository) FindPokedexes () []model.Pokedex {
	if pokedexes, found := pokedexRepository.cache.Get("pokedexes"); found {
		return pokedexes.([]model.Pokedex)
	} else {
		var pokedexes []model.Pokedex
		_ = pokedexRepository.db.Preload("Region").Find(&pokedexes)
		pokedexRepository.cache.Set("pokedexes", pokedexes, cache.DefaultExpiration)
		return pokedexes
	}
}

func (pokedexRepository PokedexRepository) FindPokedexByID(id int) model.Pokedex {
	if pokedex, found := pokedexRepository.cache.Get(fmt.Sprintf("pokedex|%d", id)); found {
		return pokedex.(model.Pokedex)
	} else {
		var pokedex model.Pokedex
		_ = pokedexRepository.db.Where(model.Pokedex{PokedexID: id}).Preload(clause.Associations).Preload("Region.Subregions").Preload("Entries.Pokemon").Preload("Entries.Pokemon.Types").Preload("Entries.Pokemon.Physique").Preload("Entries.Pokemon.Abilities").Preload("Entries.Pokemon.Statistics").Preload("Entries.Pokemon.Training").Preload("Entries.Pokemon.Images").Take(&pokedex)
		pokedexRepository.cache.Set(fmt.Sprintf("pokedex|%d", id), pokedex, cache.DefaultExpiration)
		return pokedex
	}
}
