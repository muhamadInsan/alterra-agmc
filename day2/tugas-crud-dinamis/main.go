package main

import (
	"day2/praktek6-crud-dinamis/config"
	"day2/praktek6-crud-dinamis/routes"
)

func main() {

	config.InitDB()
	e := routes.New()

	e.Logger.Fatal(e.Start(":8080"))
}
