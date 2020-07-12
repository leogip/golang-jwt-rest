package main

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"

	"github.com/leogip/golang-jwt-rest/database"
	"github.com/leogip/golang-jwt-rest/logger"
	"github.com/leogip/golang-jwt-rest/router"
)

var (
	log = logger.New("main")
)


func main() {
	godotenv.Load()

	if err := database.Connect(os.Getenv("DATABASE_URL")); err != nil {
		log.Fatal(err)
	}

	routes := router.NewRouter()

	originsOk := handlers.AllowedOrigins([]string{"*"})
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	log.Info("Server Is Running At ", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":" + os.Getenv("PORT"), handlers.CORS(headersOk, originsOk, methodsOk)(routes)))
}