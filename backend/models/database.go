package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

const IMAGES_DIR = "./images"

var All_Players Players

func Setup_DB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./players.db")
	if err != nil {
		return nil, err
	}
	create_table_sql := `
	CREATE TABLE IF NOT EXISTS players (
		id INTEGER PRIMARY KEY,
		url TEXT NOT NULL,
		elo INTEGER NOT NULL DEFAULT 1400,
		k_factor INTEGER NOT NULL DEFAULT 40,
		rounds INTEGER NOT NULL DEFAULT 0
	);
	`
	_, err = db.Exec(create_table_sql)
	if err != nil {
		return nil, err
	}

	fmt.Println("Database created succussfully")
	return db, nil
}

func (players *Players) Load_DB() {
	rows, err := players.DB.Query("SELECT id, url, elo, k_factor, rounds FROM players")
	if err != nil {
		fmt.Printf("Error loading database %v\n", err)
		return
	}
	defer rows.Close()
	// var player_list []Image
	for rows.Next() {
		var image Image
		err := rows.Scan(&image.ID, &image.URL, &image.ELO, &image.K_FACTOR, &image.ROUNDS)
		if err != nil {
			fmt.Printf("Error scanning row %v", err)
			continue
		}
		players.Images = append(players.Images, image)
	}

	fmt.Println("Loaded the DB")

	// players.Images = player_list
}

func (players *Players) Add_Player(image Image) {
	var exists bool
	err := players.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM players WHERE url = ?)", image.URL).Scan(&exists)

	if err != nil {
		fmt.Printf("Error checking existence of player %v\n", err)
		return
	}

	if exists {
		fmt.Printf("Player %s already exists\n", image.URL)
		return
	}

	_, err = players.DB.Exec("INSERT INTO players (id, url, elo, k_factor, rounds) VALUES (?, ?, ?, ?, ?)", image.ID, image.URL, image.ELO, image.K_FACTOR, image.ROUNDS)
	if err != nil {
		fmt.Printf("Error adding player %v", err)
		return
	}
	fmt.Printf("Added player into DB %s\n", image.URL)
	players.Images = append(players.Images, image)
}

func (players *Players) Update_ELO(image Image) {
	_, err := players.DB.Exec("UPDATE players SET elo = ? WHERE id = ?", image.ELO, image.ID)
	if err != nil {
		fmt.Printf("Error updating ELO %v", err)
		return
	}
}

func (players *Players) Update_Rounds(image Image) {
	_, err := players.DB.Exec("UPDATE players SET rounds = rounds + ? WHERE id = ?", 1, image.ID)
	if err != nil {
		fmt.Printf("Error updating ROUNDS %v", err)
		return
	}
}

func (players *Players) GetImagesList() error {
	fmt.Println("Getting images YAY")

	// Create images directory if it doesn't exist
	if _, err := os.Stat(IMAGES_DIR); os.IsNotExist(err) {
		log.Printf("Creating images directory: %s", IMAGES_DIR)
		if err := os.MkdirAll(IMAGES_DIR, 0755); err != nil {
			return err
		}
	}

	// Walk through the images directory
	err := filepath.Walk(IMAGES_DIR, func(path string, info os.FileInfo, err error) error {
		fmt.Printf("Path: %s\n", path)
		if err != nil {
			fmt.Printf("Error walking through this image %v\n", err)
			return err
		}
		// Skip directories
		if info.IsDir() {
			return nil
		}
		// Check if the file is an image based on extension
		ext := strings.ToLower(filepath.Ext(path))
		if ext == ".jpg" || ext == ".jpeg" || ext == ".png" || ext == ".gif" || ext == ".webp" {
			// Get just the filename without the directory path
			relPath, err := filepath.Rel(IMAGES_DIR, path)
			if err != nil {
				return err
			}
			image := Image{
				ID:       len(players.Images),
				URL:      relPath,
				ELO:      1400,
				K_FACTOR: 40,
				ROUNDS:   0,
			}
			// players.Images = append(players.Images, image)

			players.Add_Player(image)

			fmt.Printf("Added player %s\n", image.URL)
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking through images %v\n", err)
	} else {
		fmt.Println("Walked through the images directory no problemo")
	}

	fmt.Println(players.Images)

	return err
}
