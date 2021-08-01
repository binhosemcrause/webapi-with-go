package main

import (
	"github.com/binhosemcrause/webapi-with-go/db"
	"github.com/binhosemcrause/webapi-with-go/server"
)

func main() {
	db.ConectDataBase()

	server := server.NewServer()

	server.Run()

}
