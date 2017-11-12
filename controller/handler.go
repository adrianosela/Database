package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/adrianosela/Database/table"
)

type CreateTablePayload struct {
	Name       string `json:"name"`
	PrimaryKey string `json:"key"`
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

	if pl.Name == "" || pl.PrimaryKey == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Missing mandatory field \"name\" or \"key\""))
		return
	}

	t, err := table.NewTable(pl.Name, pl.PrimaryKey)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Could not write table to file system: %s", err)))
		return
	}

	c.Lock()
	c.Cache[pl.Name] = t
	c.Config.PRIMap[pl.Name] = pl.PrimaryKey //remember to solve writing to the config
	c.Unlock()

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Table \"%s\" added to database with PRI \"%s\"", pl.Name, pl.PrimaryKey)))
	return
}
