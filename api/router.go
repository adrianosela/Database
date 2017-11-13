package api

import (
	"github.com/adrianosela/Database/controller"
	"github.com/gorilla/mux"
)

//GetDatabaseRouter returns the API's router
func GetDatabaseRouter(ctrl *controller.DBController) *mux.Router {

	router := mux.NewRouter()

	//table operations
	router.Methods("GET").Path("/tables").HandlerFunc(ctrl.GetTablesHandler)
	router.Methods("POST").Path("/table").HandlerFunc(ctrl.CreateTableHandler)
	router.Methods("GET").Path("/table/{table_name}").HandlerFunc(ctrl.GetTableHandler)
	router.Methods("DELETE").Path("/table/{table_name}").HandlerFunc(ctrl.DeleteTableHandler)

	//item operations
	router.Methods("POST").Path("/item/{table_name}").HandlerFunc(ctrl.AddItemHandler)
	router.Methods("DELETE").Path("/item/{table_name}/{id}").HandlerFunc(ctrl.DeleteItemHandler)

	return router
}
