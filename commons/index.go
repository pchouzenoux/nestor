package commons

import (
	"io/ioutil"
	"os/exec"
)

func ReadFile(filepath string) string {
	data, err := ioutil.ReadFile(filepath)

	if err != nil {
		panic(err)
	}

	return string(data)
}

func ToFile(filepath string, filecontent []byte) {
	err := ioutil.WriteFile(filepath, filecontent, 0644)
	if err != nil {
		panic(err)
	}
}

func AppendToFile(filepath string, filecontent []byte) {
	err := ioutil.WriteFile(filepath, filecontent, 0644)
	if err != nil {
		panic(err)
	}
}

func ExecShellCmd(command string, args ...string) string {
	cmd := exec.Command(command, args...)

	out, err := cmd.CombinedOutput()

	if err != nil {
		panic(err)
	}

	return string(out)
}
