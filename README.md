# goalfred #

goalfred is a helper package to create workflows for Alfredapp.

**Documentation:** [![GoDoc](https://godoc.org/github.com/BenchR267/goalfred?status.svg)](https://godoc.org/github.com/BenchR267/goalfred)  
**Build Status:** [![Build Status](https://travis-ci.org/BenchR267/goalfred.svg?branch=master)](https://travis-ci.org/BenchR267/goalfred)  

## Get ##

```bash
go get -u github.com/BenchR267/goalfred
```

## Usage ##

```go
import "github.com/BenchR267/goalfred"
```

Create a new Response:

```go
res := NewResponse()
```

Construct your data like you want it to appear in the script filter output.
You can then either use the given Item struct to add elements to the output or provide your custom types:

Using Item:
```go
func main() {
	res := goalfred.NewResponse()

	item := goalfred.Item {
			Title: "aTitle",
			Subtitle: "aSubtitle",
			Arg: "https://www.example.com",
	}
	item.Mod.Alt = goalfred.ModContent {
			Arg: "https://www.google.de",
			Subtitle: "Open Google!",
  }

	res.AddItem(item)

	res.Print()
}
```

Using a custom type like Link:
```go
type Link struct {
	Name  string
	Link1 string
	Link2 string
}

func (l Link) Item() *goalfred.Item {
	item := goalfred.Item {
			Title: l.Name,
			Arg: l.Link1,
	}

	item.Mod.Cmd = goalfred.ModContent {
			Arg: l.Link2,
			Subtitle: "Something special!",
  }
	return &item
}

func main() {
	res := goalfred.NewResponse()

	link := Link{
		Name: "Google",
		Link1: "https://www.google.com",
		Link2: "https://www.google.de/search?q=hello+world",
	}

	res.AddItem(link)

	res.Print()
}
```

At the end you have to call the Print() function on the Response instance to print the elements to stdout:
```go
res.Print()
```

# Customization

Each Item has many properties, most of them are optional. To get more information about them, take a look at the official documentation at Alfred: https://www.alfredapp.com/help/workflows/inputs/script-filter/json/

## License ##

This library is distributed under the MIT-style license found in the [LICENSE](./LICENSE)
file.
