package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"sync"

	"github.com/adrianosela/Database/table"
)

type Controller struct {
	sync.RWMutex //inherit read/write lock behavior
	Config       ControllerConfig
	Cache        map[string]*table.Table
}

type ControllerConfig struct {
	ConfigFilename string
	PRIMap         map[string]string
}

func NewController(config ControllerConfig) *Controller {
	return &Controller{
		Config: config,
		Cache:  make(map[string]*table.Table),
	}
}

func NewControllerConfig(filename string) ControllerConfig {
	//read the file specified
	PRImapBytes, err := ioutil.ReadFile("./db/" + filename)
	if err != nil {
		log.Printf("[INFO] Controller configuration file not found, creating new as: ./db/%s", filename)
	}
	//unmarshall onto our config type
	var PRIMapObject map[string]string
	if err = json.Unmarshal(PRImapBytes, &PRIMapObject); err != nil {
		//if fails, return a controller config with an empty primary key map object
		log.Printf("[INFO] Controller configuration file corrupted, creating new as: ./db/%s", filename)
		return ControllerConfig{
			ConfigFilename: filename,
		}
	}
	//return a controller config with an existing config
	return ControllerConfig{
		ConfigFilename: filename,
	}
}

func (c *Controller) AddTable(t *table.Table) {
	c.Lock()
	defer c.Unlock()
	c.Cache[t.Name] = t
}
