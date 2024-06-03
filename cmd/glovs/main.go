package main

import (
	"log"
	"net/http"

	"github.com/durwankurnaik/glovs/internal/api/router"
	"github.com/durwankurnaik/glovs/internal/config"
	"github.com/durwankurnaik/glovs/internal/db"
	"github.com/durwankurnaik/glovs/internal/logger"
)

func main() {
	// load configuration
	config.LoadConfig()

	// initialize the logger
	logger.InitLogger()

	// connect to database, this part will be done afterwards
	db.ConnectDB()

	// setup the router
	r := router.SetupRouter()

	// start the server here
	log.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}