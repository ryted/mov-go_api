package MovieHandler

import (
	// Custom
	MovieType "github.com/ryted/mov-go_api/src/Model"
	MovieData "github.com/ryted/mov-go_api/src/data"

	// Built In
	"encoding/json"
	"net/http"

	// Third-Parties
	"github.com/gorilla/mux"
)

func GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	for _, item := range MovieData.MovieList {
		json.NewEncoder(w).Encode(item)
	}
}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, item := range MovieData.MovieList {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	http.Error(w, "Data tidak ditemukan!", http.StatusNotFound)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newMovie MovieType.Movie

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newMovie)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	MovieData.MovieList = append(MovieData.MovieList, newMovie)
	json.NewEncoder(w).Encode(newMovie)
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newUpdate MovieType.Movie

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newUpdate); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	params := mux.Vars(r)
	for index, item := range MovieData.MovieList {
		if item.ID == params["id"] {
			MovieData.MovieList[index] = newUpdate
			json.NewEncoder(w).Encode(MovieData.MovieList)
			return
		}
	}

	http.Error(w, "Data tidak ditemukan!", http.StatusNotFound)
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for index, item := range MovieData.MovieList {
		if item.ID == params["id"] {
			MovieData.MovieList = append(MovieData.MovieList[:index], MovieData.MovieList[index+1:]...)
			break
		}
	}
}
