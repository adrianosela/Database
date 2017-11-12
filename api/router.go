package api

import (
	"github.com/adrianosela/Database/table"
	"github.com/gorilla/mux"
)

//GetDatabaseRouter returns the API's router
func GetDatabaseRouter() *mux.Router {

	router := mux.NewRouter()

	router.Methods("POST").Path("/table").HandlerFunc(table.CreateTableHandler)

	return router
}
