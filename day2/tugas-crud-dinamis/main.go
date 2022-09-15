package main

import (
	"day2/tugas-crud-dinamis/config"
	"day2/tugas-crud-dinamis/routes"
)

func main() {

	config.InitDB()
	e := routes.New()

	e.Logger.Fatal(e.Start(":8080"))
}
