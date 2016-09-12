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
func NormalizedArguments() (normalizedArgs []string, err error) {
	for _, e := range Arguments() {
		var normalizedElement string
		normalizedElement, err = Normalize(e)
		if err != nil {
			return
		}
		normalizedArgs = append(normalizedArgs, normalizedElement)
	}
	return
}

// Response is the top level domain object.
// Create a new instance by calling NewResponse()
// Add items by calling AddItem on the response object
type Response struct {
	Items []Item `json:"items"`
}

// NewResponse initializes a new instance of Response
func NewResponse() *Response {
	r := new(Response)
	r.Items = []Item{}
	return r
}

// Print should be called last to output the result of the workflow to stdout.
func (r *Response) Print() {
	bytes, _ := json.Marshal(r)
	fmt.Println(string(bytes))
}

// AddItem adds a new Item to the response.
// The order in Alfred will be in the order how you add them.
func (r *Response) AddItem(item AlfredItem) *Response {
	i := item.Item()
	r.Items = append(r.Items, *i)
	return r
}
