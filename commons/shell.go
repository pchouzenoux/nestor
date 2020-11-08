package commons

import (
	"log"
	"os/exec"
	"strings"
)

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
