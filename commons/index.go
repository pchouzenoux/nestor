package commons

import (
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
)

func ReadFile(filepath string) string {
	data, err := ioutil.ReadFile(filepath)

	if err != nil {
		log.Fatal(err)
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

func ExecShellCmd(command string, args string) string {
	splitArgs := strings.Split(args, " ")

	log.Printf("> %s %s", command, args)

	cmd := exec.Command(command, splitArgs...)

	stdout, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}

	return string(stdout)
}
