package goalfred

import (
	"encoding/json"
	"fmt"
	"os"
)

// Arguments just wrappes the call to os.Args for better readability
func Arguments() []string {
	return os.Args[1:]
}

// NormalizedArguments re-normalizes the user arguments provided via Alfred.
// This isn't necessary for every workflow, specifically only when you're working with special characters.
// For more info on this topic, please refer to this thread: http://www.alfredforum.com/topic/2015-encoding-issue/
// Arguments that couldn't get normalized are not part of the return value!
func NormalizedArguments() (normalizedArgs []string, err error) {
	for _, e := range Arguments() {
		var normalizedElement string
		normalizedElement, err = Normalize(e)
		if err != nil {
			continue
		}
		normalizedArgs = append(normalizedArgs, normalizedElement)
	}
	return
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
	fmt.Printf("%s", jsonFromItems(items))
}
