package model

func (Region) TableName() string { return "regions" }
type Region struct {
	RegionID   int         `json:"id"         gorm:"column:regionID;primaryKey"`
	Name       string      `json:"name"       gorm:"column:name"`
	Subregions []Subregion `json:"subregions" gorm:"foreignKey:RegionID;references:RegionID"`
}