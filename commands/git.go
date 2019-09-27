package commands

import (
	"log"
	"os"
)

import (
	"github.com/Nehorim/nestor/git"
)

func Git(cmd string) {
	switch cmd {
	case "clean":
		log.Print("Clean git local repository")
		git.Clean()
	default:
		log.Fatal("Error: Command not supported, Try `nestor help` to get more information")
		os.Exit(1)
	}
}
