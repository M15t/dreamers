package main

import (
	"dreamers/routers"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	e := routers.New()

	// run server
	e.Logger.Fatal(e.Start(":9013"))
}
