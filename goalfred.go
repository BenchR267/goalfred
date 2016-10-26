package goalfred

import (
	"encoding/json"
	"log"
)

func init() {
	log.SetFlags(0) // don't output date and time
}

func jsonFromItems(items []Item) string {
	res := struct {
		Items []Item `json:"items"`
	}{
		Items: items,
	}
	bytes, _ := json.Marshal(res)
	return string(bytes)
}

var items = []Item{}

// Add adds the item to be ready to print
func Add(item AlfredItem) {
	items = append(items, item.Item())
}

// Print prints out the saved items
func Print() {
	log.Println(jsonFromItems(items))
}
