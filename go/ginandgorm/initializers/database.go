package initializers

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error

	// func gorm.Open(dialector gorm.Dialector, opts ...gorm.Option) (db *gorm.DB, err error)
	// Open initialize db session based on dialector
	// https://gorm.io/docs/index.html#Quick-Start
	DB, err = gorm.Open(sqlite.Open("./test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}
