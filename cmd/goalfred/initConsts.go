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
	<array>
		<dict>
			<key>config</key>
			<dict>
				<key>alfredfiltersresults</key>
				<false/>
				<key>argumenttype</key>
				<integer>1</integer>
				<key>escaping</key>
				<integer>0</integer>
				<key>keyword</key>
				<string>:keyword:</string>
				<key>queuedelaycustom</key>
				<integer>3</integer>
				<key>queuedelayimmediatelyinitially</key>
				<true/>
				<key>queuedelaymode</key>
				<integer>0</integer>
				<key>queuemode</key>
				<integer>1</integer>
				<key>runningsubtext</key>
				<string></string>
				<key>script</key>
				<string></string>
				<key>scriptargtype</key>
				<integer>1</integer>
				<key>scriptfile</key>
				<string>main</string>
				<key>subtext</key>
				<string></string>
				<key>title</key>
				<string>:name:</string>
				<key>type</key>
				<integer>8</integer>
				<key>withspace</key>
				<true/>
			</dict>
			<key>type</key>
			<string>alfred.workflow.input.scriptfilter</string>
			<key>uid</key>
			<string>6CA02D1B-BC8D-4C4C-9A7B-E229A8989B94</string>
			<key>version</key>
			<integer>2</integer>
		</dict>
	</array>
	<key>readme</key>
	<string></string>
    <key>uidata</key>
	<dict>
		<key>6CA02D1B-BC8D-4C4C-9A7B-E229A8989B94</key>
		<dict>
			<key>xpos</key>
			<integer>10</integer>
			<key>ypos</key>
			<integer>10</integer>
		</dict>
	</dict>
	<key>webaddress</key>
	<string>:website:</string>
</dict>
</plist>
`

	mainFileTemplate = `package main

import "github.com/BenchR267/goalfred"

func main() {
    defer goalfred.Print()

    for _, a := range goalfred.Arguments() {
        item := goalfred.Item {
            Title: a,
            Subtitle: "aSubtitle",
            Arg: "https://www.example.com",
        }
        item.Mod.Alt = &goalfred.ModContent {
            Arg: "https://www.google.de",
            Subtitle: "Open Google!",
        }
        goalfred.Add(item)
    }

}
`
)
