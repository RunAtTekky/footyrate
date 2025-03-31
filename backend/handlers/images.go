package handlers

import (
	"encoding/json"
	"image_compare/models"
	"net/http"
)

func Handle_imagelist(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(models.All_Players.Images)
}
