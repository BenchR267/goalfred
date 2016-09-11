package goalfred

import (
	"encoding/json"
	"fmt"
)

// ComplexArg gives you the opportunity to set variables as well that you can use later
type ComplexArg struct {
	Arg       string                 `json:"arg,omitempty"`
	Variables map[string]interface{} `json:"variables,omitempty"`
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

// AlfredItem defines that a struct is convertible to an Item
type AlfredItem interface {
	Item() *Item
}

// ItemType describes the type of an Item
type ItemType string

// NoItemType is the default item type, you could also leave it in the zero value state
var NoItemType ItemType

// FileItemType makes Alfred treat your result as a file on your system
var FileItemType ItemType = "file"

// SkipCheckItemType makes Alfred skip this check as you are certain that the files you are returning exist
var SkipCheckItemType ItemType = "file:skipcheck"

// Item stores informations about on item in the script filter
// A possible gotcha here is the `Valid` attribute, which is a pointer to a bool. This ensures it is whatever you set it and it gets included in the output if and only if you set it.
type Item struct {
	UID          string      `json:"uid,omitempty"`
	Title        string      `json:"title"`
	Subtitle     string      `json:"subtitle"`
	Arg          string      `json:"arg,omitempty"`
	Icon         *Icon       `json:"icon,omitempty"`
	Valid        *bool       `json:"valid,omitempty"`
	Autocomplete string      `json:"autocomplete,omitempty"`
	Type         ItemType    `json:"type,omitempty"`
	Mod          ModElements `json:"mods,omitempty"`
	Quicklook    string      `json:"quicklook,omitempty"`
}

func (c ComplexArg) string() string {
	b, _ := json.Marshal(struct {
		C ComplexArg `json:"alfredworkflow"`
	}{C: c})
	return string(b)
}

// SetComplexArg sets the argument of the item to a more complex one that could contain variables as well
func (i *Item) SetComplexArg(arg ComplexArg) {
	i.Arg = arg.string()
}

// Item is an AlfredItem
func (i Item) Item() *Item {
	return &i
}

// AddItem adds a new Item to the response.
// The order in Alfred will be in the order how you add them.
func (r *Response) AddItem(item AlfredItem) *Response {
	i := item.Item()
	r.Items = append(r.Items, *i)
	return r
}

// ModElements is a collection of the different modifiers for the item
// Alt will be visible when holding the alt-key
// Cmd will be visible when holding the cmd-key
type ModElements struct {
	Alt *ModContent `json:"alt,omitempty"`
	Cmd *ModContent `json:"cmd,omitempty"`
}

// ModContent holds all informations about a modifier of an Item
type ModContent struct {
	Valid    bool   `json:"valid,omitempty"`
	Arg      string `json:"arg,omitempty"`
	Subtitle string `json:"subtitle,omitempty"`
}

// SetComplexArg sets the argument of the item to a more complex one that could contain variables as well
func (m *ModContent) SetComplexArg(arg ComplexArg) {
	m.Arg = arg.string()
}

// IconType describes the two possible values for the Type of an icon
type IconType string

// From https://www.alfredapp.com/help/workflows/inputs/script-filter/json/
// By omitting the "type", Alfred will load the file path itself, for example a png.
// By using "type": "fileicon", Alfred will get the icon for the specified path.
// Finally, by using "type": "filetype", you can get the icon of a specific file, for example "path": "public.png"

// NoIconType makes Alfred load the file path itself
var NoIconType IconType

// FileIconType makes Alfred get the icon for the specified path
var FileIconType IconType = "fileicon"

// FileTypeIconType makes Alfed get the icon of a specific file
var FileTypeIconType IconType = "filetype"

// Icon holds all information about an item's icon
type Icon struct {
	Type IconType `json:"type,omitempty"`
	Path string   `json:"path,omitempty"`
}
