package api

import "github.com/gorilla/mux"

//GetDatabaseRouter returns the API's router
func GetDatabaseRouter() *mux.Router {

	router := mux.NewRouter()

	return router
}
