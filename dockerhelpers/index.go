package dockerhelpers

import (
	"nestor/commons"
	"strings"
)

func Clean() {
	result := commons.ExecShellCmd("docker", "ps -aq --no-trunc -f status=exited")

	for _, dockerImg := range strings.Split(result, "\n") {
		dockerImg = strings.Trim(dockerImg, "")
		if dockerImg == "" {
			continue
		}
		commons.ExecShellCmd("docker", "rm "+dockerImg)
	}

	commons.ExecShellCmd("docker", "image prune -f")
}
