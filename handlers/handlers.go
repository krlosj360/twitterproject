package handlers

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

func Manejadores() {
	router := mux.NewRouter()
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "3002"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
