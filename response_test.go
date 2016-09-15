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
		actualJSON := jsonFromItems(test.response.([]Item)...)
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

func TestItemsFromAlfredItems(t *testing.T) {
	emptyItems := itemsFromAlfredItems([]AlfredItem{})
	if len(emptyItems) != 0 {
		t.Error("Expected length of 0, got", len(emptyItems))
	}

	exampleLink := link{
		title:    "title",
		subtitle: "subtitle",
		arg:      "arg",
	}
	convertedArray := itemsFromAlfredItems([]AlfredItem{exampleLink})
	if len(convertedArray) != 1 {
		t.Error("Expected length of 1, got", len(convertedArray))
	}
	if convertedArray[0] != *exampleLink.Item() {
		t.Errorf("Converted %v to Item. Expected %v, got %v.", exampleLink, *exampleLink.Item(), convertedArray[0])
	}

}

// TODO: Not tested yet: Arguments(), NormalizedArguments(), Print()
