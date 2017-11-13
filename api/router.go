package api

import (
	"fmt"
	"net/http"

	"github.com/adrianosela/Database/controller"
	"github.com/gorilla/mux"
)

//GetDatabaseRouter returns the API's router
func GetDatabaseRouter(ctrl *controller.DBController) *mux.Router {

	router := mux.NewRouter()

	router.NotFoundHandler = http.HandlerFunc(CustomNotFoundHandler)

	//table operations
	router.Methods("GET").Path("/tables").HandlerFunc(ctrl.GetTablesHandler)
	router.Methods("POST").Path("/table").HandlerFunc(ctrl.CreateTableHandler)
	router.Methods("GET").Path("/table/{table_name}").HandlerFunc(ctrl.GetTableHandler)
	router.Methods("DELETE").Path("/table/{table_name}").HandlerFunc(ctrl.DeleteTableHandler)

	//item operations
	router.Methods("POST").Path("/item/{table_name}").HandlerFunc(ctrl.AddItemHandler)
	router.Methods("GET").Path("/item/{table_name}/{item_id}").HandlerFunc(ctrl.GetItemHandler)
	router.Methods("GET").Path("/items/{table_name}").HandlerFunc(ctrl.GetItemsHandler)
	router.Methods("DELETE").Path("/item/{table_name}/{item_id}").HandlerFunc(ctrl.DeleteItemHandler)

	return router
}

func CustomNotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(fmt.Sprintf("Invalid URL %s", r.URL.String())))
}
