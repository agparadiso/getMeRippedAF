package models

type Ingredient struct {
	ID           int    `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	Name         string `gorm:"not null" form:"name" json:"name"`
	Protein      int    `form:"protein" json:"protein"`
	Carbohydrate int    `form:"carbohydrate" json:"carbohydrate"`
	Fat          int    `form:"fat" json:"fat"`
}
