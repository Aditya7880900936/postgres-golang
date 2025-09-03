package router

import (
	"github.com/Aditya7880900936/postgres-golang/middlewares"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/stock/{id}", middlewares.GetStock).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/stock",middlewares.GetAllStock).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/newstock",middlewares.CreateStock).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/stock/{id}",middlewares.UpdateStock).Methods("PUT","OPTIONS")
	router.HandleFunc("/api/deletestock/{id}",middlewares.DeleteStock).Methods("DELETE","OPTIONS")

}
