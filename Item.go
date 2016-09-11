package goalfred

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

// Item is an AlfredItem
func (i Item) Item() *Item {
	return &i
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
