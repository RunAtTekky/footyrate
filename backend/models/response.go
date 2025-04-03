package models

// Response struct to send back JSON with image URLs
type Response struct {
	Player1 Player `json:"image1"`
	Player2 Player `json:"image2"`
}
