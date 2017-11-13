package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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
	log.Printf("[INFO] New Request For: %s\n", r.URL.String())

	varMap := mux.Vars(r)
	tableName := varMap["table_name"]
	id := varMap["item_id"]

	c.RLock()
	defer c.RUnlock()

	if t, okt := c.Tables[tableName]; okt {

		log.Println("[INFO] Got to A")

		t.RLock()
		defer t.RUnlock()

		if o, oko := t.Cache[id]; oko {

			log.Println("[INFO] Got to B")

			if respBytes, err := json.Marshal(o); err == nil {
				w.WriteHeader(http.StatusOK)
				w.Write(respBytes)
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Could not marshall item for response"))
			return
		}
		// w.WriteHeader(http.StatusNotFound)
		// w.Write([]byte(fmt.Sprintf("Item with id %s not found in table %s", id, tableName)))
		// return
	}
	// w.WriteHeader(http.StatusNotFound)
	// w.Write([]byte(fmt.Sprintf("Table %s not found", tableName)))
	// return

	log.Println("[INFO] Got to C")
	//get bytes from file
	fileBytes, err := ioutil.ReadFile(fmt.Sprintf("./db/%s/%s", tableName, id))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, fmt.Sprintf("Item with path tablename/id: %s/%s not found", tableName, id))
		return
	}

	log.Println("[INFO] Got to D")
	var jsonObject map[string]interface{}
	if err = json.Unmarshal(fileBytes, &jsonObject); err != nil {
		//gracefull failure of this part
		log.Printf("[ERROR] Could not unmarshall item with path  %s/%s \n", tableName, id)
	} else {
		log.Println("[INFO] Got to E")
		c.Tables[tableName].Cache[id] = jsonObject
	}

	log.Println("[INFO] Got to F")
	w.WriteHeader(http.StatusOK)
	w.Write(fileBytes)
	return
}

func (c *DBController) GetItemsHandler(w http.ResponseWriter, r *http.Request) {
	//TODO
	return
}

func (c *DBController) DeleteItemHandler(w http.ResponseWriter, r *http.Request) {
	//TODO
	return
}
