# goalfred #

goalfred is a helper package to create workflows for Alfredapp.

[![GoDoc](https://godoc.org/github.com/BenchR267/goalfred?status.svg)](https://godoc.org/github.com/BenchR267/goalfred)  [![Build Status](https://travis-ci.org/BenchR267/goalfred.svg?branch=master)](https://travis-ci.org/BenchR267/goalfred) [![Coverage Status](https://coveralls.io/repos/github/BenchR267/goalfred/badge.svg)](https://coveralls.io/github/BenchR267/goalfred)

## Get ##

```bash
go get -u github.com/BenchR267/goalfred
```

## Usage ##

```go
import "github.com/BenchR267/goalfred"
```

Construct your data like you want it to appear in the script filter output.
You can then either use the given Item struct to add elements to the output or provide your custom types:

Using Item:
```go
func main() {

	defer goalfred.Print()

	item := goalfred.Item {
			Title: "aTitle",
			Subtitle: "aSubtitle",
			Arg: "https://www.example.com",
	}
	item.Mod.Alt = goalfred.ModContent {
			Arg: "https://www.google.de",
			Subtitle: "Open Google!",
  }

	goalfred.Add(item)
}
```

Using a custom type like Link:
```go
type Link struct {
	Name  string
	Link1 string
	Link2 string
}

func (l Link) Item() goalfred.Item {
	item := goalfred.Item {
			Title: l.Name,
			Arg: l.Link1,
	}

	item.Mod.Cmd = goalfred.ModContent {
			Arg: l.Link2,
			Subtitle: "Something special!",
  }
	return item
}

func main() {

	defer goalfred.Print()

	link := Link{
		Name: "Google",
		Link1: "https://www.google.com",
		Link2: "https://www.google.de/search?q=hello+world",
	}

	goalfred.Add(link)
}
```

# Customization

Each Item has many properties, most of them are optional. To get more information about them, take a look at the [official documentation](https://www.alfredapp.com/help/workflows/inputs/script-filter/json/) at Alfred.

# Complex Arguments

Sometimes it's necessary to output more than one parameter by the workflow. For example if you want to schedule a notification like [alfred_dvb](https://github.com/kiliankoe/alfred_dvb) does. If an entry was selected there are multiple informations that should be outhanded. The first one is the time after which the notification should be triggered and the second one is the text for the notification. You could add a script that splits a given string into pieces, but with goalfred you can also add complex arguments that contain a query AND parameters.
You can achieve this by doing the following:

```Go
item := goalfred.Item{
	Title: "a title",
	Subtitle: "a subtitle",
}
item.SetComplexArg(goalfred.ComplexArg{
	Arg: "A nice query that can be used as {query} in Alfred actions",
	Variables: map[string]interface{}{
		"time": 5,
		"output": "this text will replace {var:output} in Alfred actions!",
	},
})
```

As you can already see in the code, you can specify variables as well. You can use then the variables value by writing {var:VARIABLENAME} in an Alfred action. Very handy!

## License ##

This library is distributed under the MIT-style license found in the [LICENSE](./LICENSE)
file.
