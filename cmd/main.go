package main

import (
	"github.com/gorilla/mux"
	"http-proxy-server/internal/handler"
	"log"
	"net/http"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", handler.HandleFunc)

	err := http.ListenAndServe(":8080", router)

	if err != nil {
		log.Fatal(err)
	}
	log.Println("Starting server on :8080")

}
