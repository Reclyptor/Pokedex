package model

type Pokedex []Pokemon

func (Pokemon) TableName() string { return "pokemon" }
type Pokemon struct {
	PokemonID    int        `json:"-"          gorm:"column:pokemonID;primaryKey"`
	PokedexID    int        `json:"id"         gorm:"column:pokedexID"`
	Generation   int        `json:"generation" gorm:"column:generation"`
	IsDefault    bool       `json:"isDefault"  gorm:"column:isDefault"`
	Name         string     `json:"name"       gorm:"column:name"`
	Species      string     `json:"species"    gorm:"column:species"`
	Variant      string     `json:"variant"    gorm:"column:variant"`
	Status       string     `json:"status"     gorm:"column:status"`
	Shape        string     `json:"shape"      gorm:"column:shape"`
	Types        Types      `json:"types"      gorm:"foreignKey:PokemonID"`
	Physique     Physique   `json:"physique"   gorm:"foreignKey:PokemonID"`
	Abilities    Abilities  `json:"abilities"  gorm:"foreignKey:PokemonID"`
	Statistics   Statistics `json:"statistics" gorm:"foreignKey:PokemonID"`
	Training     Training   `json:"training"   gorm:"foreignKey:PokemonID"`
	Images       []Image    `json:"images"     gorm:"foreignKey:PokemonID"`
}

func (Types) TableName() string { return "types" }
type Types struct {
	PokemonID int    `json:"-"         gorm:"column:pokemonID;primaryKey"`
	Primary   string `json:"primary"   gorm:"column:primary"`
	Secondary string `json:"secondary" gorm:"column:secondary"`
}

func (Physique) TableName() string { return "physique" }
type Physique struct {
	PokemonID int `json:"-"      gorm:"column:pokemonID;primaryKey"`
	Height    int `json:"height" gorm:"column:height"`
	Weight    int `json:"weight" gorm:"column:weight"`
}

func (Abilities) TableName() string { return "abilities" }
type Abilities struct {
	PokemonID int    `json:"-"         gorm:"column:pokemonID;primaryKey"`
	Primary   string `json:"primary"   gorm:"column:primary"`
	Secondary string `json:"secondary" gorm:"column:secondary"`
	Hidden    string `json:"hidden"    gorm:"column:hidden"`
}

func (Statistics) TableName() string { return "statistics" }
type Statistics struct {
	PokemonID      int `json:"-"              gorm:"column:pokemonID;primaryKey"`
	HP             int `json:"hp"             gorm:"column:hp"`
	Attack         int `json:"attack"         gorm:"column:attack"`
	Defense        int `json:"defense"        gorm:"column:defense"`
	SpecialAttack  int `json:"specialAttack"  gorm:"column:specialAttack"`
	SpecialDefense int `json:"specialDefense" gorm:"column:specialDefense"`
	Speed          int `json:"speed"          gorm:"column:speed"`
	Total          int `json:"total"          gorm:"column:total"`
}

func (Training) TableName() string { return "training" }
type Training struct {
	PokemonID  int    `json:"-"          gorm:"column:pokemonID;primaryKey"`
	CatchRate  int    `json:"catchRate"  gorm:"column:catchRate"`
	GrowthRate string `json:"growthRate" gorm:"column:growthRate"`
}

func (Image) TableName() string { return "images" }
type Image struct {
	PokemonID int    `json:"-"     gorm:"column:pokemonID;primaryKey"`
	Image     string `json:"image" gorm:"column:image"`
	Src       string `json:"src"   gorm:"column:src"`
}