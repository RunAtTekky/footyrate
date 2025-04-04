package main

import (
	"context"
	"fmt"
	"image_compare/models"
	"image_compare/server"
	"log"
	"net/http"

	"github.com/robfig/cron"
)

const PORT = 8080
const API = "https://footyrate.onrender.com"

func main() {

	c := cron.New()
	c.AddFunc("@every 14m", ping_server)
	log.Println("Started cron job")
	c.Start()

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

	select {}
}

func ping_server() {

	res, err := http.Get(API)

	if err != nil {
		fmt.Printf("Error getting response from server %v", err)
	}

	res.Body.Close()

	println("Pinged the server")
}
