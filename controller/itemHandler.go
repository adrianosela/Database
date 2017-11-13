package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func (c *DBController) AddItemHandler(w http.ResponseWriter, r *http.Request) {
	varMap := mux.Vars(r)
	tableName := varMap["table_name"]

	c.RLock()
	defer c.RUnlock()

	if t, ok := c.Tables[tableName]; ok {

		bodyBytes, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("Could not read request body: %s", err)))
			return
		}

		var item map[string]interface{}
		err = json.Unmarshal(bodyBytes, &item)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("Could not unmarshall malformed payload: %s", err)))
			return
		}

		err = t.AddItem(item)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("Could not write to file system: %s", err)))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Added new item to table"))
		return
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(fmt.Sprintf("No such table: \"%s\"", tableName)))
	return
}

func (c *DBController) GetItemHandler(w http.ResponseWriter, r *http.Request) {
	//TODO
	return
}

func (c *DBController) DeleteItemHandler(w http.ResponseWriter, r *http.Request) {
	//TODO
	return
}
