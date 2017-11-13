package table

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

//Table represents a new table with a cache
type Table struct {
	sync.RWMutex                        //inherit read/write lock behavior
	Name         string                 `json:"name"`
	Cache        map[string]interface{} `json:"cache,omitempty"`
}

//NewTable makes a new table by creating a new directory with the table name
func NewTable(name string) (*Table, error) {
	err := os.MkdirAll(fmt.Sprintf("./db/%s", name), 755)
	if err != nil {
		return nil, err
	}
	return &Table{
		Name:  name,
		Cache: make(map[string]interface{}),
	}, nil
}

func (t *Table) AddItem(JSONitem map[string]interface{}) error {
	if id, ok := JSONitem["id"]; ok {
		if idStr, ok := id.(string); ok {
			t.Lock()
			defer t.Unlock()
			//TODO: evict from cache if n > SOME_THRESHOLD
			t.Cache[idStr] = JSONitem
			if err := writeItem(fmt.Sprintf("./db/%s/%s", t.Name, idStr), JSONitem); err != nil {
				return fmt.Errorf("Could not write item %s to table %s on filesystem: %s", idStr, t.Name, err)
			}
			return nil
		}
		return errors.New("Mandatory JSON tag \"id\" must be a string")
	}
	return errors.New("Missing mandatory JSON tag \"id\" in JSON Object")
}

func (t *Table) DeleteItem(id string) error {
	t.Lock()
	defer t.Unlock()
	if _, ok := t.Cache[id]; ok {
		delete(t.Cache, id)
		return nil
	}
	return fmt.Errorf("Did not find object %s in table %s", id, t.Name)
}

func writeItem(path string, item map[string]interface{}) error {
	dataBytes, err := json.Marshal(item)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, dataBytes, 755)
}
