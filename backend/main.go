package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/rand/v2"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

type Image struct {
	ID  int    `json:"id"`
	URL string `json:"url"`
	ELO int    `json:"elo"`
}

// Response struct to send back JSON with image URLs
type Response struct {
	Image1 Image `json:"image1"`
	Image2 Image `json:"image2"`
}

type Result struct {
	Winner_ID int `json:"winner_ID"`
	Loser_ID  int `json:"loser_ID"`
}

var Images []Image

// Configuration settings
const (
	PORT        = 8080
	IMAGES_DIR  = "./images" // Directory containing your images
	IMAGES_URL  = "/images/" // URL path to access images
	RANDOM_PATH = "/api/random-images"
)

type Server struct {
	Router *chi.Mux
}

func CreateServer() *Server {
	server := &Server{
		Router: chi.NewRouter(),
	}
	return server
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello"))
}

func (server *Server) MountHandlers() {
	server.Router.Get("/", hello)

	apiRouter := chi.NewRouter()

	apiRouter.Group(func(r chi.Router) {
		r.Get("/random-images", handle_random)
		r.Get("/images", handle_imagelist)
		r.Post("/result", handle_result)
	})

	server.Router.Handle(IMAGES_URL+"*", http.StripPrefix(IMAGES_URL, http.FileServer(http.Dir(IMAGES_DIR))))

	server.Router.Mount("/api", apiRouter)
}

func handle_result(w http.ResponseWriter, r *http.Request) {

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	w.WriteHeader(http.StatusOK)

	result := new(Result)

	if err := json.NewDecoder(r.Body).Decode(result); err != nil {
		fmt.Printf("Error decoding result")
		return
	}

	// fmt.Printf("Winner %d \n Loser %d\n\n", result.Winner_ID, result.Loser_ID)

	compute_result(result)

	// w.Write([]byte(result.Winner))

	json.NewEncoder(w).Encode(result)
}

func compute_result(result *Result) {

	// result.Winner
	winner_ID := result.Winner_ID
	loser_ID := result.Loser_ID

	update_ELO(winner_ID, loser_ID)

}

func update_ELO(winner_ID int, loser_ID int) {
	// TODO - Write logic for updating elo

	winner_ELO := Images[winner_ID].ELO
	loser_ELO := Images[loser_ID].ELO

	var difference_ELO float32 = float32(winner_ELO) - float32(loser_ELO)

	expected := 1 / (math.Pow(10, float64(difference_ELO/400)) + 1)

	K := 20.0

	ELO_change := K * (1 - expected)

	winner_ELO += int(ELO_change)
	loser_ELO -= int(ELO_change)

	Images[winner_ID].ELO = winner_ELO
	Images[loser_ID].ELO = loser_ELO

	fmt.Printf("ID: %d ELO: %d\n", winner_ID, Images[winner_ID].ELO)
	fmt.Printf("ID: %d ELO: %d\n\n", loser_ID, Images[loser_ID].ELO)
}

func handle_random(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	w.WriteHeader(http.StatusOK)

	// images, err := getImagesList()

	// if err != nil {
	// 	fmt.Println("Error getting random images")
	// 	return
	// }

	baseURL := getBaseURL(r)

	Image1_idx, Image2_idx := get_two_images(&Images)

	img1 := Images[Image1_idx]
	img2 := Images[Image2_idx]

	img1.URL = baseURL + IMAGES_URL + img1.URL
	img2.URL = baseURL + IMAGES_URL + img2.URL

	// img1 := Image{
	// 	ID:  Image1_idx,
	// 	URL: baseURL + IMAGES_URL + images[Image1_idx],
	// 	ELO: 1400,
	// }
	// img2 := Image{
	// 	ID:  Image2_idx,
	// 	URL: baseURL + IMAGES_URL + images[Image2_idx],
	// 	ELO: 1400,
	// }

	response := Response{
		Image1: img1,
		Image2: img2,
	}

	json.NewEncoder(w).Encode(response)
}

func get_two_images(images *[]Image) (int, int) {

	var image1 int
	var image2 int

	n := len(*images)

	image1 = rand.IntN(n)
	image2 = rand.IntN(n)

	for image1 == image2 {
		image2 = rand.IntN(n)
	}

	return image1, image2

}

func handle_imagelist(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	w.WriteHeader(http.StatusOK)

	// images, err := getImagesList()

	// if err != nil {
	// 	fmt.Println("Error getting images")
	// 	w.Write([]byte("LOL Error getting the images"))
	// 	return
	// }

	json.NewEncoder(w).Encode(Images)
}

func main() {

	getImagesList()

	server := CreateServer()

	// CORS Middleware Configuration
	server.Router.Use(cors.Handler(cors.Options{
		// AllowedOrigins specifies the allowed origins
		// Use "*" to allow all origins (not recommended for production)
		AllowedOrigins: []string{
			"*",
			// "http://localhost:3000",    // React app
			// "https://yourfrontend.com", // Production frontend
			// "http://localhost:8080",    // Vue/Angular dev server
		},

		// AllowedMethods specifies the allowed HTTP methods
		AllowedMethods: []string{
			"GET", "POST", "PUT", "DELETE",
			"OPTIONS", "PATCH", "HEAD",
		},

		// AllowedHeaders specifies the allowed headers
		AllowedHeaders: []string{
			"Accept",
			"Authorization",
			"Content-Type",
			"X-CSRF-Token",
			"X-Requested-With",
		},

		// ExposedHeaders specifies headers that can be accessed by the client
		ExposedHeaders: []string{
			"Link",
			"X-Total-Count",
		},

		// AllowCredentials allows cookies and authentication
		AllowCredentials: true,

		// MaxAge specifies how long preflight request can be cached
		MaxAge: 300, // 5 minutes
	}))

	server.MountHandlers()

	// Start the server
	log.Printf("Server started on http://localhost:8080")
	http.ListenAndServe(":8080", server.Router)
}

// getImagesList returns a list of image filenames from the images directory
func getImagesList() ([]string, error) {
	var images []string

	var curr_ID int = 1

	// Create images directory if it doesn't exist
	if _, err := os.Stat(IMAGES_DIR); os.IsNotExist(err) {
		log.Printf("Creating images directory: %s", IMAGES_DIR)
		if err := os.MkdirAll(IMAGES_DIR, 0755); err != nil {
			return nil, err
		}
	}

	// Walk through the images directory
	err := filepath.Walk(IMAGES_DIR, func(path string, info os.FileInfo, err error) error {
		if err != nil {
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
			images = append(images, relPath)
			image := Image{
				ID:  len(Images),
				URL: relPath,
				ELO: 1400,
			}
			Images = append(Images, image)
			curr_ID++
		}
		return nil
	})

	return images, err
}

// getBaseURL constructs the base URL from the request
func getBaseURL(r *http.Request) string {
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	return fmt.Sprintf("%s://%s", scheme, r.Host)
}
