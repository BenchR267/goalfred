package goalfred

import (
	"bytes"
	"log"
	"os"
	"strings"
	"testing"
)

// thanks to http://stackoverflow.com/a/26806093
func captureOutput(f func()) string {
	var buf bytes.Buffer
	log.SetFlags(0)
	log.SetOutput(&buf)
	f()
	log.SetOutput(os.Stdout)
	return strings.TrimSpace(buf.String())
}

type responseTestPair struct {
	response interface{}
	json     string
}

var responseTests = []responseTestPair{
	{
		response: []Item{},
		json:     `{"items":[]}`,
	},
	{
		response: []Item{
			Item{
				Title:    "foo",
				Subtitle: "bar",
			},
		},
		json: `{"items":[{"title":"foo","subtitle":"bar","mods":{}}]}`,
	},
}

func TestJsonFromItems(t *testing.T) {
	for _, test := range responseTests {
		actualJSON := jsonFromItems(test.response.([]Item))
		if test.json != string(actualJSON) {
			t.Error("Expected", test.json, "got", string(actualJSON))
		}
	}
}

type link struct {
	title    string
	subtitle string
	arg      string
}

func (l link) Item() *Item {
	return &Item{
		Title:    l.title,
		Subtitle: l.subtitle,
		Arg:      l.arg,
	}
}

func TestAdd(t *testing.T) {
	items = []Item{}
	i := Item{}
	Add(i)
	if len(items) != 1 {
		t.Error("items should be length 1 after adding one element.")
	}

	if items[0] != i {
		t.Errorf("item was not correctly added. Expected first item to be %v, got %v.", i, items[0])
	}
}

func TestPrint(t *testing.T) {
	items = []Item{}
	output := captureOutput(func() {
		Print()
	})

	if output != "{\"items\":[]}" {
		t.Error("Expected output with empty items. Got: ", output)
	}

	i := Item{
		Title:    "a title",
		Subtitle: "a subtitle",
	}
	Add(i)

	output = captureOutput(func() {
		Print()
	})
	if output != "{\"items\":[{\"title\":\"a title\",\"subtitle\":\"a subtitle\",\"mods\":{}}]}" {
		t.Error("Expected output with empty items. Got: ", output)
	}
}

func TestRerun(t *testing.T) {
	items = []Item{}
	Rerun(2)
	output := captureOutput(func() {
		Print()
	})

	if output != "{\"rerun\":2,\"items\":[]}" {
		t.Error("Expected output with empty items and rerun set to 2. Got: ", output)
	}
}

// TODO: Not tested yet: Arguments(), NormalizedArguments()
