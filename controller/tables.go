package controller

import "github.com/adrianosela/Database/table"

func (c *DBController) AddTable(t *table.Table) {
	c.Lock()
	defer c.Unlock()
	c.Tables[t.Name] = t
}

func (c *DBController) DeleteTable(t *table.Table) {
	c.Lock()
	defer c.Unlock()
	delete(c.Tables, t.Name)
}
