package goalfred

import (
	"fmt"
	"os"
)

// IsDebug returns if the current environment has 'alfred_debug' set.
func IsDebug() bool {
	alfredDebug := os.Getenv("alfred_debug")
	return alfredDebug != ""
}

// Log logs text to the debug panel. Only logs when IsDebug is true, so as not
// to interfere with the normal output.
func Log(text string) {
	if IsDebug() {
		fmt.Println(text)
	}
}
