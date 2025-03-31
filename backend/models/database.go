package models

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

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

	return db, nil
}

func (players *Players) Load_DB() {

	rows, err := players.DB.Query("SELECT id, url, elo, k_factor, rounds FROM players")

	if err != nil {
		fmt.Printf("Error loading database %v", err)
		return
	}

	defer rows.Close()

	var player_list []Image

	for rows.Next() {
		var image Image

		err := rows.Scan(&image.ID, &image.URL, &image.ELO, &image.K_FACTOR, &image.ROUNDS)

		if err != nil {
			fmt.Printf("Error scanning row %v", err)
			continue
		}

		player_list = append(player_list, image)
	}

	players.Images = player_list
}

func (players *Players) Add_Player(image *Image) {

	_, err := players.DB.Exec("INSERT INTO players (id, url, elo, k_factor, rounds) VALUES (?, ?, ?, ?, ?)", image.ID, image.URL, image.ELO, image.K_FACTOR, image.ROUNDS)

	if err != nil {
		fmt.Printf("Error adding player %v", err)
		return
	}

	players.Images = append(players.Images, *image)

}

func (players *Players) Update_ELO(image *Image) {

	_, err := players.DB.Exec("UPDATE players SET elo = ? WHERE id = ?", image.ELO, image.ID)

	if err != nil {
		fmt.Printf("Error updating ELO %v", err)
		return
	}

}

func (players *Players) Update_Rounds(image *Image) {

	_, err := players.DB.Exec("UPDATE players SET rounds = rounds + ? WHERE id = ?", 1, image.ID)

	if err != nil {
		fmt.Printf("Error updating ROUNDS %v", err)
		return
	}

}
