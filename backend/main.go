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

	defer db.Close()

	// All_Players := &models.Players{
	// 	Images: []models.Image{},
	// 	DB:     db,
	// }

	handlers.All_Players.Images = []models.Image{
		{
			ID:       0,
			URL:      "Cris",
			K_FACTOR: 40,
			ELO:      1600,
			ROUNDS:   5,
		},
	}
	handlers.All_Players.DB = db

	// Load the images
	err = handlers.GetImagesList()

	if err != nil {
		log.Fatalf("Error occurred in getting images")
		return
	}

	// Load the database
	handlers.All_Players.Load_DB()

	// Server
	srv := server.CreateServer()

	server.Configure_CORS(srv)

	srv.MountHandlers()

	// Start the server
	log.Printf("Server started on http://localhost:%d", PORT)
	http.ListenAndServe(fmt.Sprintf(":%d", PORT), srv.Router)
}
