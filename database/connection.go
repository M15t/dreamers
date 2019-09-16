package database

import (
	"dreamers/models"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

func InitDB(filepath string) *gorm.DB {
	db, err = gorm.Open("sqlite3", filepath)
	if err != nil {
		fmt.Println(err)
	}
	//defer database.Close() // I consider that will this run one or run on each request

	// migrate as well
	if !(db.HasTable(&models.Diary{})) || !(db.HasTable(&models.Person{})) {
		db.AutoMigrate(&models.Diary{}, &models.Person{})
	}

	/// old method and test
	//for _, model := range []interface{}{
	//	Diary{}, Person{},
	//} {
	//	if err := database.AutoMigrate(model).Error; err != nil {
	//		fmt.Println(err)
	//	} else {
	//		fmt.Println("Auto migrating", reflect.TypeOf(model).Name(), "...")
	//	}
	//}

	// here we check for any database errors then exit
	if err != nil {
		panic(err)
	}

	// if we don't get any errors but somehow still don't get a database connection
	// we exit as well
	if db == nil {
		panic("database nil")
	}
	return db
}
