package model

func (Pokedex) TableName() string { return "pokedex" }
type Pokedex struct {
	PokedexID int             `json:"id"      gorm:"column:pokedexID;primaryKey"`
	Name      string          `json:"name"    gorm:"column:name"`
	RegionID  int             `json:"-"       gorm:"column:regionID"`
	Region    Region          `json:"region"  gorm:"foreignKey:RegionID;references:RegionID"`
	Entries   []PokedexPokemon`json:"entries" gorm:"foreignKey:PokedexID;references:PokedexID"`
}