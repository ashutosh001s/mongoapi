package controller

import (
	"encoding/json"
	"log"
	"net/http"

	dbs "github.com/ashutosh001s/mongoapi/db"
	"github.com/ashutosh001s/mongoapi/model"
	"github.com/gorilla/mux"
)

// Actual controller file

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allMovies := dbs.GetAllMovies()
	if err := json.NewEncoder(w).Encode(allMovies); err!=nil{
		log.Fatal(err)
	}
}

func CreateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix
	err:=json.NewDecoder(r.Body).Decode(&movie)
	if err!=nil {
		log.Fatal(err)
	}
	dbs.InsertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)

}

func MarkAsWatched(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	dbs.UpdateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteMovie(w http.ResponseWriter, r * http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params:= mux.Vars(r)

	dbs.DeleteOneMovie(params["id"])
	json.NewEncoder(w).Encode("deleted one record")
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := dbs.DeleteAllMovie()
	json.NewEncoder(w).Encode(count)
}

