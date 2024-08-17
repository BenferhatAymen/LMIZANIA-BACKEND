package main

import (
	"lmizania/config"
	"lmizania/database"
	"lmizania/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	config.LoadENV()
	database.MongoClient = database.MongoDBInstance()
	database.RedisClient = database.RedisDBInstance()

	router := mux.NewRouter()

	routes.AuthRoutes(router)
	routes.TransactionRoutes(router)

	http.ListenAndServe(":8080", router)

}
