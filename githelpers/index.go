package githelpers

import (
	"fmt"
	"log"
	"nestor/commons"
	"strings"
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
	log.Printf(commons.ExecShellCmd("git", "remote prune origin"))

	result := commons.ExecShellCmd("git", "branch")

	branches := strings.Split(result, "\n")
	for i, branch := range branches {
		branch = strings.Trim(branch, " ")
		if branch == "" || strings.HasPrefix(branch, "*") || branchIsExcluded(branch) {
			continue
		}
		fmt.Println(i, branch)
		log.Printf(commons.ExecShellCmd("git", "branch -d "+branch))
	}
}
