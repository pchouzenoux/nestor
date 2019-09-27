package commands


import (
	"log"
	"os"
)

import (
	"github.com/Nehorim/nestor/docker"
)

func Docker(cmd string) {
	switch cmd {
	case "clean":
		log.Print("Clean docker images")
		docker.Clean()
	default:
		log.Fatal("Error: Command not supported, Try `nestor help` to get more information")
		os.Exit(1)
	}
}
