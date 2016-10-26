package goalfred

import (
	"bytes"
	"log"
	"os"
	"testing"
)

func TestIsDebug(t *testing.T) {
	os.Unsetenv("alfred_debug")
	if IsDebug() {
		t.Errorf("Expected return value of IsDebug to be false if env var is not set.")
	}

	os.Setenv("alfred_debug", "true")
	if !IsDebug() {
		t.Errorf("Expected return value of IsDebug to be true if env var is set.")
	}
}

// thanks to http://stackoverflow.com/a/26806093
func captureOutput(f func()) string {
	var buf bytes.Buffer
	log.SetFlags(0)
	log.SetOutput(&buf)
	f()
	log.SetOutput(os.Stdin)
	return buf.String()
}

func TestLog(t *testing.T) {
	os.Unsetenv("alfred_debug")
	emptyOutput := captureOutput(func() {
		Log("some log")
	})
	if emptyOutput != "" {
		t.Errorf("Expected output to be empty since debug is not turned on. Got: %s", emptyOutput)
	}

	os.Setenv("alfred_debug", "true")
	someOutput := captureOutput(func() {
		Log("some log")
	})
	if someOutput != "some log\n" {
		t.Errorf("Expected output to be 'some log' since debug is turned on. Got: %s", someOutput)
	}
}
