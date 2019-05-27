package commons

import (
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
)

func ReadFile(filepath string) (string, error) {
	data, err := ioutil.ReadFile(filepath)

	if err != nil {
		return "", err
	}

	return string(data), nil
}

func ToFile(filepath string, filecontent []byte) error {
	err := ioutil.WriteFile(filepath, filecontent, 0644)
	return err
}

func AppendToFile(filepath string, filecontent []byte) error {
	err := ioutil.WriteFile(filepath, filecontent, 0644)
	return err
}

func ExecShellCmd(command string, args string) (string, error) {
	splitArgs := strings.Split(args, " ")

	log.Printf("> %s %s", command, args)

	cmd := exec.Command(command, splitArgs...)

	stdout, err := cmd.CombinedOutput()
	if err != nil {
		return string(stdout), err
	}

	return string(stdout), nil
}
