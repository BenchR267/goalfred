package goalfred

import (
	"encoding/json"
	"testing"
)

type complexargTestPair struct {
	complexarg *ComplexArg
	json       string
}

var complexargTests = []complexargTestPair{
	{
		complexarg: &ComplexArg{},
		json:       `{"alfredworkflow":{}}`,
	},
	{
		complexarg: &ComplexArg{
			Arg: "arg",
			Variables: map[string]interface{}{
				"foo": "bar",
			},
		},
		json: `{"alfredworkflow":{"arg":"arg","variables":{"foo":"bar"}}}`,
	},
}

func TestComplexArgJSON(t *testing.T) {
	for _, test := range complexargTests {
		actualJSON := test.complexarg.string()
		if test.json != actualJSON {
			t.Error("Expected", test.json, "got", actualJSON)
		}
	}
}

func TestItemSetComplexArg(t *testing.T) {
	item := Item{}
	item.SetComplexArg(ComplexArg{
		Arg: "arg",
		Variables: map[string]interface{}{
			"foo": "bar",
		},
	})
	actualJSON, err := json.Marshal(item)
	expected := `{"title":"","subtitle":"","arg":"{\"alfredworkflow\":{\"arg\":\"arg\",\"variables\":{\"foo\":\"bar\"}}}","mods":{}}`
	if expected != string(actualJSON) || err != nil {
		t.Error("Expected", expected, "got", string(actualJSON))
	}
}

func TestModContentSetComplexArg(t *testing.T) {
	mod := ModContent{}
	mod.SetComplexArg(ComplexArg{
		Arg: "arg",
		Variables: map[string]interface{}{
			"foo": "bar",
		},
	})
	actualJSON, err := json.Marshal(mod)
	expected := `{"arg":"{\"alfredworkflow\":{\"arg\":\"arg\",\"variables\":{\"foo\":\"bar\"}}}"}`
	if expected != string(actualJSON) || err != nil {
		t.Error("Expected", expected, "got", string(actualJSON))
	}
}
