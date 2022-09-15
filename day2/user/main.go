package main

import (
	"admc-day2-user/config"
	"admc-day2-user/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	e.Logger.Fatal(e.Start(":8000"))
}
