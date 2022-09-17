package main

import "day2/tugas-crud-statis/routes"

func main() {

	e := routes.New()

	e.Logger.Fatal(e.Start(":8080"))
}
