package main

import (
	"fmt"
	"image_compare/handlers"
	"image_compare/models"
	"image_compare/server"
	"log"
	"net/http"
)

const PORT = 8080

func main() {

	// Setup database
	db, err := models.Setup_DB()

	if err != nil {
		log.Fatalf("Error setting database %v", err)
		return
	}

	db.Close()

	// Load the database
	handlers.All_Players.Load_DB()

	// Load the images
	err = handlers.GetImagesList()

	if err != nil {
		log.Fatalf("Error occurred in getting images")
		return
	}

	// Server
	srv := server.CreateServer()

	server.Configure_CORS(srv)

	srv.MountHandlers()

	// Start the server
	log.Printf("Server started on http://localhost:%d", PORT)
	http.ListenAndServe(fmt.Sprintf(":%d", PORT), srv.Router)
}
