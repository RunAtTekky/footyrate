package models

// Response struct to send back JSON with image URLs
type Response struct {
	Image1 Image `json:"image1"`
	Image2 Image `json:"image2"`
}
