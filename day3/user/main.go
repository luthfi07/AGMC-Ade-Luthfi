package main

import (
	"admc-day2-user/config"
	m "admc-day2-user/middlewares"
	"admc-day2-user/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	m.LogMiddlewares(e)
	e.Logger.Fatal(e.Start(":8000"))
}
