package controller

import (
	"sync"

	"github.com/adrianosela/Database/table"
)

type DBController struct {
	sync.RWMutex //inherit read/write lock behavior
	Tables       map[string]*table.Table
}

func NewDBController() *DBController {
	return &DBController{
		Tables: make(map[string]*table.Table),
	}
}
