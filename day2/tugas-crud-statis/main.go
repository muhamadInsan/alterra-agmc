package main

import "day2/praktek5-homework/routes"

func main() {
	e := routes.Init()

	e.Logger.Fatal(e.Start(":8080"))
}
