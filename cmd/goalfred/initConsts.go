package main

const (
	infoPlistTemplate = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>bundleid</key>
	<string>:bundle:</string>
	<key>category</key>
	<string>:category:</string>
	<key>connections</key>
	<dict/>
	<key>createdby</key>
	<string>:author:</string>
	<key>description</key>
	<string>:description:</string>
	<key>disabled</key>
	<false/>
	<key>name</key>
	<string>:name:</string>
	<key>objects</key>
	<array/>
	<key>readme</key>
	<string></string>
	<key>uidata</key>
	<dict/>
	<key>webaddress</key>
	<string>:website:</string>
</dict>
</plist>
`

	mainFileTemplate = `package main

import "github.com/BenchR267/goalfred"

func main() {
    defer goalfred.Print()

    item := goalfred.Item {
            Title: "aTitle",
            Subtitle: "aSubtitle",
            Arg: "https://www.example.com",
    }
    item.Mod.Alt = &goalfred.ModContent {
            Arg: "https://www.google.de",
            Subtitle: "Open Google!",
  }

    goalfred.Add(item)
}
`
)
