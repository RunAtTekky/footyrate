package models

type Image struct {
	ID       int    `json:"id"`
	URL      string `json:"url"`
	ELO      int    `json:"elo"`
	K_FACTOR int    `json:"k_factor"`
	ROUNDS   int    `json:"rounds"`
}
