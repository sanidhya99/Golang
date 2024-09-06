package router

import (
	"github.com/gorilla/mux"
	"github.com/sanidhya99/mongoapi/controllers"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/v1/get/movies", controllers.GetMyAllMovies).Methods("GET")
	router.HandleFunc("/v1/post/movies", controllers.CreateMovie).Methods("POST")
	router.HandleFunc("/v1/edit/movies/{id}", controllers.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/v1/delete/movies/{id}", controllers.DeleteOneMovie).Methods("DELETE")
	router.HandleFunc("/v1/delete/movies", controllers.DeleteAllMovie).Methods("DELETE")
	return router
}
