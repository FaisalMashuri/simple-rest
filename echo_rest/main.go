package main

import (
	"echo_rest/db"
	"echo_rest/routes"
)

func main() {
	db.NewDB()
	e := routes.Init()
	e.Logger.Fatal(e.Start(":1323"))
}
