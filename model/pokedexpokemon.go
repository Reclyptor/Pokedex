package model

func (PokedexPokemon) TableName() string { return "pokedex_pokemon" }
type PokedexPokemon struct {
	PokedexID     int     `json:"-"           gorm:"column:pokedexID"`
	PokedexNumber int     `json:"id"          gorm:"column:pokedexNumber"`
	SubregionID   int     `json:"subregionID" gorm:"column:subregionID"`
	PokemonID     int     `json:"-"           gorm:"column:pokemonID"`
	Pokemon       Pokemon `json:"pokemon"     gorm:"foreignKey:PokemonID;references:PokemonID"`
}