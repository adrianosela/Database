package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/adrianosela/Database/table"
	"github.com/gorilla/mux"
)

type CreateTablePayload struct {
	Name string `json:"name"`
}

func (c *DBController) CreateTableHandler(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Could not read request body: %s", err)))
		return
	}

	var pl CreateTablePayload
	if err = json.Unmarshal(bodyBytes, &pl); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Could not unmarshall malformed payload: %s", err)))
		return
	}

	if pl.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Missing mandatory field \"name\""))
		return
	}

	t, err := table.NewTable(pl.Name)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Could not write table to file system: %s", err)))
		return
	}

	c.AddTable(t)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Table \"%s\" added to database", pl.Name)))
	return
}

func (c *DBController) GetTableHandler(w http.ResponseWriter, r *http.Request) {
	varMap := mux.Vars(r)
	tableName := varMap["table_name"]

	//get files in table directory
	files, err := ioutil.ReadDir(fmt.Sprintf("./db/%s", tableName))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, fmt.Sprintf("Table \"%s\" not found", tableName))
		return
	}

	response := &struct {
		Objects []string `json:"objects"`
	}{
		Objects: []string{},
	}

	for _, f := range files {
		response.Objects = append(response.Objects, f.Name())
	}

	respBytes, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not marshall response"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(respBytes)
	return
}

func (c *DBController) GetTablesHandler(w http.ResponseWriter, r *http.Request) {
	files, err := ioutil.ReadDir("./db")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, []byte("Could not read tables directory"))
		return
	}

	response := &struct {
		Tables []string `json:"tables"`
	}{
		Tables: []string{},
	}

	for _, f := range files {
		response.Tables = append(response.Tables, f.Name())
	}

	respBytes, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not marshall response"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(respBytes)
	return
}

func (c *DBController) DeleteTableHandler(w http.ResponseWriter, r *http.Request) {
	varMap := mux.Vars(r)
	tableName := varMap["table_name"]

	//Check can read table directory
	_, err := ioutil.ReadDir(fmt.Sprintf("./db/%s", tableName))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, fmt.Sprintf("Table \"%s\" not found", tableName))
		return
	}

	if err = os.RemoveAll(fmt.Sprintf("./db/%s", tableName)); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, fmt.Sprintf("Table \"%s\" could not be deleted", tableName))
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, fmt.Sprintf("Table \"%s\" successfully deleted", tableName))
	return
}
