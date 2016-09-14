package goalfred

import (
	"io/ioutil"
	"os/exec"
)

// Normalize fixes problems with string encoding regarding the usage of special characters in Alfred.
// For more info on this topic, please refer to this thread: http://www.alfredforum.com/topic/2015-encoding-issue/
func Normalize(input string) (output string, err error) {
	iconv := exec.Command("iconv", "-f", "UTF-8")
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
