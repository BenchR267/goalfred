package goalfred

import "testing"

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

// TODO: Not tested yet: Arguments(), NormalizedArguments(), Print()
