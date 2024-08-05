package routes

import (
	"context"
	"lmizania/controllers"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var mongoClient *mongo.Client

func init() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("env load error", err)

	}

	log.Println("env file loaded")

	mongoClient, err = mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("MONGO_URI")))

	if err != nil {
		log.Fatal("connection error", err)
	}
	mongoClient.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("ping failed", err)
	}
	log.Println("mongo connected ")

}

func UserRoutes(router *mux.Router) {

	coll := mongoClient.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("USER_COLLECTION"))
	userService := controllers.UserService{MongoCollection: coll}

	router.HandleFunc("/login", userService.LoginUser).Methods(http.MethodPost)
	router.HandleFunc("/register", userService.RegisterUser).Methods(http.MethodPost)

}
