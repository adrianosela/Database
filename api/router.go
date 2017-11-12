package api

import (
	"github.com/adrianosela/Database/table"
	"github.com/gorilla/mux"
)

//GetKeystoreRouter returns the API's router
func GetKeystoreRouter() *mux.Router {

	router := mux.NewRouter()

	router.Methods("POST").Path("/table").HandlerFunc(table.CreateTableHandler)

	return router
}
