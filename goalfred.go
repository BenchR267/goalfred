package goalfred

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	Items []Item `json:"items"`
}

func NewResponse() *Response {
	return new(Response)
}

func (r *Response) AddItem(item *Item) *Response {
	r.Items = append(r.Items, *item)
	return r
}

func (r *Response) Print() {
	bytes, _ := json.Marshal(r)
	fmt.Println(string(bytes))
}

type Item struct {
	UID          string       `json:"uid,omitempty"`
	Title        string       `json:"title"`
	Subtitle     string       `json:"subtitle"`
	Arg          string       `json:"arg,omitempty"`
	Icon         string       `json:"icon,omitempty"`
	Valid        bool         `json:"valid"`
	Autocomplete string       `json:"autocomplete,omitempty"`
	Type         string       `json:"type,omitempty"`
	Mod          *ModElements `json:"mod,omitempty"`
	Quicklook    string       `json:"quicklook,omitempty"`
}

type ModElements struct {
	Alt *ModContent `json:"alt,omitempty"`
	Cmd *ModContent `json:"cmd,omitempty"`
}

func NewModElement(arg string, subtitle string) *ModContent {
	m := new(ModContent)
	m.Arg = arg
	m.Subtitle = subtitle
	m.Valid = true
	return m
}

type ModContent struct {
	Valid    bool   `json:"valid"`
	Arg      string `json:"arg,omitempty"`
	Subtitle string `json:"subtitle,omitempty"`
}

func NewItem(uid string, title string, subtitle string, arg string) *Item {
	item := new(Item)
	item.Title = title
	item.Subtitle = subtitle
	item.Arg = arg
	item.Valid = true
	item.Mod = new(ModElements)
	return item
}
