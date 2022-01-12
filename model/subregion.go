package model

func (Subregion) TableName() string { return "subregions" }
type Subregion struct {
	SubregionID int    `json:"id"   gorm:"column:subregionID;primaryKey"`
	Name        string `json:"name" gorm:"column:name"`
	RegionID    int    `json:"-"    gorm:"column:regionID"`
}