package docker

import (
	"strings"

	"github.com/Nehorim/nestor/commons"
)

func Clean() {
	result, _ := commons.ExecShellCmd("docker", "ps -aq --no-trunc -f status=exited")

	for _, dockerImg := range strings.Split(result, "\n") {
		dockerImg = strings.Trim(dockerImg, "")
		if dockerImg == "" {
			continue
		}
		commons.ExecShellCmd("docker", "rm "+dockerImg)
	}

	commons.ExecShellCmd("docker", "system prune -f")
	commons.ExecShellCmd("docker", "volume prune -f")
}
