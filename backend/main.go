package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

// Response struct to send back JSON with image URLs
type Response struct {
	Image1 string `json:"image1"`
	Image2 string `json:"image2"`
}

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

type Result struct {
	Winner string `json:"winner"`
	Loser  string `json:"loser"`
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

	json.NewEncoder(w).Encode(result)
}

func handle_random(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	w.WriteHeader(http.StatusOK)

	images, err := getImagesList()

	if err != nil {
		fmt.Println("Error getting random images")
		return
	}

	baseURL := getBaseURL(r)

	Image1, Image2 := get_two_images(&images)

	IMAGE1_URL := baseURL + IMAGES_URL + images[Image1]
	IMAGE2_URL := baseURL + IMAGES_URL + images[Image2]

	response := Response{
		Image1: IMAGE1_URL,
		Image2: IMAGE2_URL,
	}

	json.NewEncoder(w).Encode(response)
}

func get_two_images(images *[]string) (int, int) {

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

	images, err := getImagesList()

	if err != nil {
		fmt.Println("Error getting images")
		w.Write([]byte("LOL Error getting the images"))
		return
	}

	json.NewEncoder(w).Encode(images)
}

func main() {

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
