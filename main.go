package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/adrianosela/Database/api"
)

func main() {

	//check the database directory is in place
	err := checkPreconditions()
	if err != nil {
		log.Fatalf("Could not read db directory. %s", err)
	}

	router := api.GetDatabaseRouter()

	log.Println("[INFO] Listening on http://localhost:80")

	err = http.ListenAndServe(":80", router)
	if err != nil {
		log.Fatal("Could not start server: ", err)
	}

}

func checkPreconditions() error {
	_, err := ioutil.ReadDir("./db")
	return err
}
