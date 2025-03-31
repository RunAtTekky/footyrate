package server

import (
	"image_compare/handlers"
	"net/http"

	"github.com/go-chi/chi"
)

// Configuration settings
const (
	PORT        = 8080
	IMAGES_DIR  = "../images" // Directory containing your images
	IMAGES_URL  = "/images/"  // URL path to access images
	RANDOM_PATH = "/api/random-images"
)

func (server *Server) MountHandlers() {
	server.Router.Get("/", handlers.Greeting)

	apiRouter := chi.NewRouter()

	apiRouter.Group(func(r chi.Router) {
		r.Get("/random-images", handlers.Handle_random)
		r.Get("/images", handlers.Handle_imagelist)
		r.Post("/result", handlers.Handle_result)
	})

	server.Router.Handle(IMAGES_URL+"*", http.StripPrefix(IMAGES_URL, http.FileServer(http.Dir(IMAGES_DIR))))

	server.Router.Mount("/api", apiRouter)
}
