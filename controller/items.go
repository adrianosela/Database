package controller

import "fmt"

func (c *DBController) AddItemToTable(tablename string, item map[string]interface{}) error {
	c.RLock()
	defer c.RUnlock()
	if table, ok := c.Tables[tablename]; ok {
		err := table.AddItem(item)
		if err != nil {
			table.RLock()
			defer table.RUnlock()
			return fmt.Errorf("Could not add item to table \"%s\", %s", table.Name, err)
		}
		return nil
	}

	return fmt.Errorf("Could not find table \"%s\"", tablename)
}
