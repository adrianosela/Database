package table

import (
	"errors"
	"fmt"
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
			return nil
		}
		return errors.New("Mandatory JSON tag \"id\" must be a string")
	}
	return errors.New("Missing mandatory JSON tag \"id\" in JSON Object")
}
