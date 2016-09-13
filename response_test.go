package goalfred

import (
	"encoding/json"
	"testing"
)

type responseTestPair struct {
	response *Response
	json     string
}

var responseTests = []responseTestPair{
	{
		response: &Response{},
		json:     `{"items":null}`, // should it really be null? o.O
	},
	{
		response: &Response{
			Items: []Item{},
		},
		json: `{"items":[]}`,
	},
	{
		response: &Response{
			Items: []Item{
				Item{
					Title:    "foo",
					Subtitle: "bar",
				},
			},
		},
		json: `{"items":[{"title":"foo","subtitle":"bar","mods":{}}]}`,
	},
}

func TestResponseJSON(t *testing.T) {
	for _, test := range responseTests {
		actualJSON, err := json.Marshal(test.response)
		if test.json != string(actualJSON) || err != nil {
			t.Error("Expected", test.json, "got", string(actualJSON))
		}
	}
}

func TestAddItem(t *testing.T) {
	res := NewResponse()
	if len(res.Items) > 0 {
		t.Error("Response Items Count should start at 0, not", len(res.Items))
	}
	res.AddItem(Item{Title: "foobar"})
	if len(res.Items) != 1 {
		t.Error("Response Items Count should be 1 after adding a single item, not", len(res.Items))
	}
}

// TODO: Not tested yet: Arguments(), NormalizedArguments(), Print()
