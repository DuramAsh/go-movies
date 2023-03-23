package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var db = GetDB()

func getMovies(w http.ResponseWriter, r *http.Request) {
	var movies []Movie
	db.Find(&movies)
	json.NewEncoder(w).Encode(movies)
}

func RunServer() {
	http.HandleFunc("/getMovies", getMovies)
	fmt.Println("Server is listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
