package models

type Food struct {
	ID          int          `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	Name        string       `gorm:"not null" form:"name" json:"name"`
	Ingredients []Ingredient `gorm:"not null" form:"ingredients" json:"ingredients"`
}
