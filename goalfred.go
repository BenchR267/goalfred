package goalfred

import (
	"encoding/json"
	"log"
	"os"
)

func init() {
	log.SetFlags(0) // don't output date and time
	log.SetOutput(os.Stdout)
}

func jsonFromItems(items []Item) string {
	res := struct {
		Rerun     int                    `json:"rerun,omitempty"`
		Variables map[string]interface{} `json:"variables,omitempty"`
		Items     []Item                 `json:"items"`
	}{
		Rerun: rerun,
		Items: items,
	}
	if len(variables) > 0 {
		res.Variables = variables
	}
	bytes, _ := json.Marshal(res)
	return string(bytes)
}

var rerun int
var variables = make(map[string]interface{})
var items = []Item{}

// Add adds the item to be ready to print
func Add(item AlfredItem) {
	items = append(items, item.Item())
}

// Rerun sets the interval after how many seconds the workflow should run again
func Rerun(seconds int) {
	rerun = seconds
}

// SetVariable sets the value of a workflow wide variable which is passed as env var to the workflow
func SetVariable(key string, value interface{}) {
	variables[key] = value
}

// Print prints out the saved items
func Print() {
	log.Println(jsonFromItems(items))
}
