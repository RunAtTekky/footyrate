package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi/v5"
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

// func enableCors(w *http.ResponseWriter) {
// 	(*w).Header().Set("Access-Control-Allow-Origin", "*")
// 	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
// 	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
// }

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
		// r.Get("/random", get_random)
		// r.Get("/imagelist", get_image_list)
		r.Get("/random-images", get_random)
		r.Get("/images", get_image_list)
	})

	server.Router.Mount("/api", apiRouter)
}

func get_random(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	images, err := getImagesList()

	if err != nil {
		fmt.Println("Error getting random images")
		return
	}

	response := Response{
		Image1: images[0],
		Image2: images[1],
	}

	json.NewEncoder(w).Encode(response)

	// w.Write([]byte("Random"))
}

func get_image_list(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	images, err := getImagesList()

	if err != nil {
		fmt.Println("Error getting images")
		return
	}

	baseURL := getBaseURL(r)

	for _, image := range images {
		final_URL := baseURL + IMAGES_URL + image
		fmt.Println(final_URL)
	}

	w.Write([]byte("Image list"))
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
	// Set up the API endpoint for random images
	// http.HandleFunc(RANDOM_PATH, func(w http.ResponseWriter, r *http.Request) {
	// 	// Enable CORS for all requests
	// 	enableCors(&w)

	// 	// Handle preflight OPTIONS requests
	// 	if r.Method == "OPTIONS" {
	// 		w.WriteHeader(http.StatusOK)
	// 		return
	// 	}

	// 	randomImagesHandler(w, r)
	// })

	// // Set up static file server for images with CORS support
	// http.HandleFunc(IMAGES_URL, func(w http.ResponseWriter, r *http.Request) {
	// 	// Enable CORS for image requests
	// 	enableCors(&w)

	// 	// Handle preflight OPTIONS requests
	// 	if r.Method == "OPTIONS" {
	// 		w.WriteHeader(http.StatusOK)
	// 		return
	// 	}

	// 	// Strip the prefix and serve the files
	// 	fileServer := http.StripPrefix(IMAGES_URL, http.FileServer(http.Dir(IMAGES_DIR)))
	// 	fileServer.ServeHTTP(w, r)
	// })

	// Start the server
	log.Printf("Server started on http://localhost:8080")
	http.ListenAndServe(":8080", server.Router)
}

// func randomImagesHandler(w http.ResponseWriter, r *http.Request) {
// 	// Set content type for JSON response
// 	w.Header().Set("Content-Type", "application/json")

// 	// Get list of available images
// 	images, err := getImagesList()
// 	if err != nil {
// 		http.Error(w, "Failed to get images list", http.StatusInternalServerError)
// 		log.Printf("Error getting images list: %v", err)
// 		return
// 	}

// 	// If we have fewer than 2 images, return an error
// 	if len(images) < 2 {
// 		http.Error(w, "Not enough images available", http.StatusInternalServerError)
// 		return
// 	}

// 	// Select two random images
// 	image1, image2 := selectTwoRandomImages(images)

// 	// Create the image URLs
// 	baseURL := getBaseURL(r)
// 	image1URL := baseURL + IMAGES_URL + image1
// 	image2URL := baseURL + IMAGES_URL + image2

// 	// Create and send the response
// 	response := Response{
// 		Image1: image1URL,
// 		Image2: image2URL,
// 	}

// 	// Marshal the response to JSON
// 	jsonResponse, err := json.Marshal(response)
// 	if err != nil {
// 		http.Error(w, "Failed to create JSON response", http.StatusInternalServerError)
// 		log.Printf("Error marshaling JSON: %v", err)
// 		return
// 	}

// 	// Send the response
// 	w.Write(jsonResponse)
// }

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

// selectTwoRandomImages selects two different random images from the given list
// func selectTwoRandomImages(images []string) (string, string) {
// 	// If we have exactly 2 images, return both
// 	if len(images) == 2 {
// 		return images[0], images[1]
// 	}

// 	// Get first random image
// 	idx1 := rand.Intn(len(images))
// 	image1 := images[idx1]

// 	// Get second random image (must be different from the first)
// 	idx2 := idx1
// 	for idx2 == idx1 {
// 		idx2 = rand.Intn(len(images))
// 	}
// 	image2 := images[idx2]

// 	return image1, image2
// }

// getBaseURL constructs the base URL from the request
func getBaseURL(r *http.Request) string {
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	return fmt.Sprintf("%s://%s", scheme, r.Host)
}
