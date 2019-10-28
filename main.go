package main

import (
	"modb/modb"
	"modb/server"
)

func main(){
	db := modb.NewMoDB()
	server.Serve(db)
}