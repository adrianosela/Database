package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	log.Println("[INFO] Listening on http://localhost:80")

	err := http.ListenAndServe(":80", router)
	if err != nil {
		log.Fatal("Could not start server: ", err)
	}

}
