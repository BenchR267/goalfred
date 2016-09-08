# Usage

```Go
func main() {
	res := NewResponse()

	item := NewItem("aTitle", "aSubtitle", "https://www.example.com")
	item.Mod.Alt = NewModElement("https://www.google.de", "Open Google!")

	res.AddItem(item)

	res.Print()
}
```

You can also provide your own types by implementing AlfredItem:

```Go
type Link struct {
	Name  string
	Link1 string
	Link2 string
}

func (l Link) Item() *Item {
	item := NewItem(l.Name, l.Name, "", l.Link1)
	item.Mod.Cmd = NewModElement(l.Link2, "Something special")
	return &item
}

func main() {
	res := NewResponse()

	link := Link{Name: "Google", Link1: "https://www.google.com", Link2: "https://www.google.de/search?q=hello+world"}

	res.AddItem(link)

	res.Print()
}
```