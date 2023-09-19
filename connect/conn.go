package connect

import (
	models "github.com/Asad2730/anstAPI/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Connect() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database ")
	}
	db.AutoMigrate(&models.Measurement{})
	Db = db
}
