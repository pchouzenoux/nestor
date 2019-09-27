package main


import (
	"os"
)

import (
	"github.com/Nehorim/nestor/commands"
)

func main() {
	// output := flag.String("ouput", "./", "Output directory. Default: ./")
	// flag.Parse()

	if len(os.Args) != 3 {
		commands.Help()
		return
	}
	switch os.Args[1] {
	case "init":
		commands.InitProject(os.Args[2])
	case "git":
		commands.Git(os.Args[2])
	case "docker":
		commands.Docker(os.Args[2])
	default:
		commands.Help()
	}
}
