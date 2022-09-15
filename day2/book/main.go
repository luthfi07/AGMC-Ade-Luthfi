package main

import (
	"admc-day2-book/routes"
)

func main() {
	e := routes.New()
	e.Logger.Fatal(e.Start(":8000"))
}
