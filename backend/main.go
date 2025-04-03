package main

import (
	"context"
	"fmt"
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

	defer func() {
		if err := db.Client().Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	models.All_Players.Player_List = []models.Player{}
	models.All_Players.DB = db

	// Get all the players from the database
	err = models.All_Players.Get_Players()
	if err != nil {
		log.Fatalf("Error occurred in loading database")
		return
	}

	// Load the images
	err = models.All_Players.GetImagesList()

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
