package main

import (
	"github.com/sergiolucena1/database"
	"github.com/sergiolucena1/routes/server"
)

func main(){
	database.StartDB()

	server := server.NewServer()

	server.Run()
}
