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
		Rerun int    `json:"rerun,omitempty"`
		Items []Item `json:"items"`
	}{
		Rerun: rerun,
		Items: items,
	}
	bytes, _ := json.Marshal(res)
	return string(bytes)
}

var rerun int
var items = []Item{}

// Add adds the item to be ready to print
func Add(item AlfredItem) {
	items = append(items, item.Item())
}

// Rerun sets the interval after how many seconds the workflow should run again
func Rerun(seconds int) {
	rerun = seconds
}

// Print prints out the saved items
func Print() {
	log.Println(jsonFromItems(items))
}
