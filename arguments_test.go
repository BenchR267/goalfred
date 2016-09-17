package goalfred

import (
	"os"
	"testing"
)

func TestNormalize(t *testing.T) {
	input := "äöü"
	expected := "äöü"

	normalized, err := Normalize(input)
	if normalized != expected || err != nil {
		t.Error("Expected", expected, "got", normalized)
	}
}

func TestArguments(t *testing.T) {
	args := Arguments()
	if len(args) != len(os.Args[1:]) {
		t.Errorf("Arguments length is not correct. Expected %v, got %v.", args, os.Args[1:])
	}

	for i, e := range args {
		if os.Args[i+1] != e {
			t.Errorf("Argument at index %v is not correct. Expected %v, got %v.", i, os.Args[i], e)
		}
	}
}

func TestNormalizedArguments(t *testing.T) {
	args, err := NormalizedArguments()

	if err != nil {
		t.Errorf("err should be nil, but got %v", err)
	}

	if len(args) != len(os.Args[1:]) {
		t.Errorf("Arguments length is not correct. Expected %v, got %v.", args, os.Args[1:])
	}

	for i, e := range args {
		if os.Args[i+1] != e {
			t.Errorf("Argument at index %v is not correct. Expected %v, got %v.", i, os.Args[i], e)
		}
	}
}
