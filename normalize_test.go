package goalfred

import "testing"

func TestNormalize(t *testing.T) {
	input := "äöü"
	expected := "äöü"

	normalized, err := Normalize(input)
	if normalized != expected || err != nil {
		t.Error("Expected", expected, "got", normalized)
	}
}
