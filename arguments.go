package goalfred

import (
	"io/ioutil"
	"os"
	"os/exec"
)

// Normalize fixes problems with string encoding regarding the usage of special characters in Alfred.
// For more info on this topic, please refer to this thread: http://www.alfredforum.com/topic/2015-encoding-issue/
func Normalize(input string) (output string, err error) {
	iconv := exec.Command("iconv", "-f", "UTF8-MAC")
	iconvIn, err := iconv.StdinPipe()
	iconvOut, err := iconv.StdoutPipe()

	iconv.Start()
	iconvIn.Write([]byte(input))
	iconvIn.Close()

	iconvOutput, err := ioutil.ReadAll(iconvOut)
	iconv.Wait()

	output = string(iconvOutput)

	return
}

// Arguments just wrappes the call to os.Args for better readability
func Arguments() []string {
	return os.Args[1:]
}

// NormalizedArguments re-normalizes the user arguments provided via Alfred.
// This isn't necessary for every workflow, specifically only when you're working with special characters.
// For more info on this topic, please refer to this thread: http://www.alfredforum.com/topic/2015-encoding-issue/
// Arguments that couldn't get normalized are not part of the return value!
func NormalizedArguments() (normalizedArgs []string, err error) {
	for _, e := range Arguments() {
		var normalizedElement string
		normalizedElement, err = Normalize(e)
		if err != nil {
			continue
		}
		normalizedArgs = append(normalizedArgs, normalizedElement)
	}
	return
}
