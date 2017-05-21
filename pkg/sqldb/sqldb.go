package sqldb

import (
	"github.com/agparadiso/getMeRippedAF/pkg/models"
	"github.com/jinzhu/gorm"
)

func InitDb() *gorm.DB {
	// Openning file
	db, err := gorm.Open("sqlite3", "./data.db")
	db.LogMode(true)
	// Error
	if err != nil {
		panic(err)
	}

	if !db.HasTable(&models.Ingredient{}) {
		db.CreateTable(&models.Ingredient{})
		db.CreateTable(&models.Food{})
		//db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&models.Ingredient{})
	}

	return db
}
