package handlers

import "net/http"

func Greeting(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	data := `Welcome to FootyRate

These are the endpoints you can use
/random-images : get two random images
/images : get all the images list
/result : send votes`
	w.Write([]byte(data))
}
