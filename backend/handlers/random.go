package handlers

import (
	"encoding/json"
	"fmt"
	"image_compare/models"
	"math/rand/v2"
	"net/http"
)

const IMAGES_URL = "/images/" // URL path to access images

func Handle_random(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	w.WriteHeader(http.StatusOK)

	baseURL := getBaseURL(r)

	fmt.Printf("Length of All_Players.Images %d\n", len(All_Players.Images))
	Image1_idx, Image2_idx := get_two_random_indexes(len(All_Players.Images))

	img1 := All_Players.Images[Image1_idx]
	img2 := All_Players.Images[Image2_idx]

	img1.URL = baseURL + IMAGES_URL + img1.URL
	img2.URL = baseURL + IMAGES_URL + img2.URL

	fmt.Printf("Image 1: %s", img1.URL)
	fmt.Printf("Image 2: %s", img2.URL)

	response := models.Response{
		Image1: img1,
		Image2: img2,
	}

	json.NewEncoder(w).Encode(response)
}

func get_two_random_indexes(n int) (int, int) {

	var idx1 int
	var idx2 int

	idx1 = rand.IntN(n)
	idx2 = rand.IntN(n)

	for idx1 == idx2 {
		idx2 = rand.IntN(n)
	}

	return idx1, idx2

}

// getBaseURL constructs the base URL from the request
func getBaseURL(r *http.Request) string {
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	return fmt.Sprintf("%s://%s", scheme, r.Host)
}
