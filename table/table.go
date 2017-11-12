package table

import (
	"fmt"
	"os"
)

type Table struct {
	Name  string                 `json:"name"`
	Items map[string]interface{} `json:"items,omitempty"`
}

func NewTable(name string) (*Table, error) {
	err := os.MkdirAll(fmt.Sprintf("./db/%s", name), 755)
	if err != nil {
		return nil, err
	}

	return &Table{
		Name:  name,
		Items: make(map[string]interface{}),
	}, nil
}
