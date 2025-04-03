package models

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const IMAGES_DIR = "./images"

var All_Players Players

func Setup_DB() (*mongo.Database, error) {
	uri := os.Getenv("MONGODB_URI")
	docs := "www.mongodb.com/docs/drivers/go/current/"
	if uri == "" {
		log.Fatal("Set your 'MONGODB_URI' environment variable. " +
			"See: " + docs +
			"usage-examples/#environment-variable")
	}

	ctx := context.Background()
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		panic(err)
	}

	DB := client.Database("footyrate")
	return DB, nil
}

func (players *Players) Add_Player(player *Player) {

	coll := players.DB.Collection("players")

	cursor, err := coll.Find(context.TODO(), bson.D{{Key: "url", Value: player.URL}})

	if err != nil {
		fmt.Printf("Error getting players with URL %v\n", err)
		return
	}

	var result []Player

	cursor.All(context.TODO(), &result)

	if len(result) != 0 {
		fmt.Printf("Player with URL: %s already present\n", player.URL)
		return
	}

	_, err = coll.InsertOne(context.TODO(), player)

	if err != nil {
		log.Printf("Error adding player to DB %v\n", err)
		return
	}

	log.Printf("Added a new player with _id %s\n", player.ID)

	All_Players.Player_List = append(All_Players.Player_List, *player)
}

func (players *Players) Update_Rounds(player *Player) {

	coll := players.DB.Collection("players")

	filter := bson.D{{Key: "url", Value: player.URL}}

	update := bson.D{{Key: "$set", Value: bson.D{{Key: "rounds", Value: player.ROUNDS}}}}

	_, err := coll.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		fmt.Printf("Error updating player's rounds %v\n", err)
	}
}

func (players *Players) Update_ELO(player *Player) {

	coll := players.DB.Collection("players")

	filter := bson.D{{Key: "url", Value: player.URL}}

	update := bson.D{{Key: "$set", Value: bson.D{{Key: "elo", Value: player.ELO}}}}

	_, err := coll.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		fmt.Printf("Error updating player's elo %v\n", err)
	}
}

func (players *Players) Get_Players() error {

	coll := players.DB.Collection("players")

	cursor, err := coll.Find(context.TODO(), bson.D{})

	if err != nil {
		fmt.Printf("Error getting players %v", err)
		return err
	}

	var result []Player

	err = cursor.All(context.TODO(), &result)

	if err != nil {
		fmt.Printf("Error decoding into player_list %v\n", err)
		return err
	}

	All_Players.Player_List = result

	return nil
}

// func Setup_DB() (*sql.DB, error) {
// 	db, err := sql.Open("sqlite3", "./players.db")
// 	if err != nil {
// 		return nil, err
// 	}
// 	create_table_sql := `
// 	CREATE TABLE IF NOT EXISTS players (
// 		id INTEGER PRIMARY KEY,
// 		url TEXT NOT NULL,
// 		elo INTEGER NOT NULL DEFAULT 1400,
// 		k_factor INTEGER NOT NULL DEFAULT 40,
// 		rounds INTEGER NOT NULL DEFAULT 0
// 	);
// 	`
// 	_, err = db.Exec(create_table_sql)
// 	if err != nil {
// 		return nil, err
// 	}

// 	fmt.Println("Database created succussfully")
// 	return db, nil
// }

// func (players *Players) Load_DB() {
// 	rows, err := players.DB.Query("SELECT id, url, elo, k_factor, rounds FROM players")
// 	if err != nil {
// 		fmt.Printf("Error loading database %v\n", err)
// 		return
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var player Player
// 		err := rows.Scan(&player.ID, &player.URL, &player.ELO, &player.K_FACTOR, &player.ROUNDS)
// 		if err != nil {
// 			fmt.Printf("Error scanning row %v", err)
// 			continue
// 		}
// 		players.Player_List = append(players.Player_List, player)
// 	}

// 	fmt.Println("Loaded the DB")
// }

// func (players *Players) Add_Player(player Player) {
// 	var exists bool
// 	err := players.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM players WHERE url = ?)", player.URL).Scan(&exists)

// 	if err != nil {
// 		fmt.Printf("Error checking existence of player %v\n", err)
// 		return
// 	}

// 	if exists {
// 		fmt.Printf("Player %s already exists\n", player.URL)
// 		return
// 	}

// 	result, err := players.DB.Exec("INSERT INTO players (url, elo, k_factor, rounds) VALUES (?, ?, ?, ?)", player.URL, player.ELO, player.K_FACTOR, player.ROUNDS)
// 	if err != nil {
// 		fmt.Printf("Error adding player %v", err)
// 		return
// 	}

// 	id, err := result.LastInsertId()

// 	if err != nil {
// 		fmt.Printf("Error getting last ID %v\n", err)
// 		return
// 	}
// 	player.ID = int(id)

// 	players.Player_List = append(players.Player_List, player)
// }

// func (players *Players) Update_ELO(player Player) {
// 	_, err := players.DB.Exec("UPDATE players SET elo = ? WHERE id = ?", player.ELO, player.ID)
// 	if err != nil {
// 		fmt.Printf("Error updating ELO %v", err)
// 		return
// 	}
// }

// func (players *Players) Update_Rounds(player Player) {
// 	_, err := players.DB.Exec("UPDATE players SET rounds = rounds + ? WHERE id = ?", 1, player.ID)
// 	if err != nil {
// 		fmt.Printf("Error updating ROUNDS %v", err)
// 		return
// 	}
// }

func (players *Players) GetImagesList() error {

	// Create images directory if it doesn't exist
	if _, err := os.Stat(IMAGES_DIR); os.IsNotExist(err) {
		log.Printf("Creating images directory: %s", IMAGES_DIR)
		if err := os.MkdirAll(IMAGES_DIR, 0755); err != nil {
			return err
		}
	}

	// Walk through the images directory
	err := filepath.Walk(IMAGES_DIR, func(path string, info os.FileInfo, err error) error {
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
			image := Player{
				ID:       primitive.NewObjectID(),
				URL:      relPath,
				ELO:      1400,
				K_FACTOR: 40,
				ROUNDS:   0,
			}

			players.Add_Player(&image)
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking through images %v\n", err)
	} else {
		fmt.Println("Walked through the images directory no problemo")
	}

	fmt.Printf("Total images: %d\n", len(players.Player_List))

	return err
}
