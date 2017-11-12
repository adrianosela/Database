package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/adrianosela/Database/table"
)

type CreateTablePayload struct {
	Name string `json:"name"`
}

func (c *Controller) CreateTableHandler(w http.ResponseWriter, r *http.Request) {
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

func (c *Controller) GetTableHandler(w http.ResponseWriter, r *http.Request) {
	//TODO
	return
}
