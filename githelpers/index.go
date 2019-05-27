package githelpers

import (
	"log"
	"strings"
)

import (
	"github.com/Nehorim/nestor/commons"
)

var excludedBranches = []string{"dev", "prod", "master"}

func branchIsExcluded(branch string) bool {
	for _, excludedBranch := range excludedBranches {
		if excludedBranch == branch {
			return true
		}
	}
	return false
}

func Clean() {
	result, err := commons.ExecShellCmd("git", "branch")

	if err != nil {
		panic(err)
	}

	branches := strings.Split(result, "\n")
	for _, branch := range branches {
		branch = strings.Trim(branch, " ")
		if branch == "" || strings.HasPrefix(branch, "*") || branchIsExcluded(branch) {
			continue
		}
		stdout, _ := commons.ExecShellCmd("git", "branch -d "+branch)
		log.Printf(stdout)
	}
	stdout, _ := commons.ExecShellCmd("git", "remote prune origin")
	log.Printf(stdout)
}
