package main

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"sample_echo/routers"
)

func main() {
	e := routers.New()

	// run server
	e.Logger.Fatal(e.Start(":9013"))
}
