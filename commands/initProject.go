package commands

import (
	"log"
	"os"

	"github.com/Nehorim/nestor/python"
	"github.com/Nehorim/nestor/typescript"
)

func InitProject(lang string) {
	switch lang {
	case "python":
		log.Print("Init python virtualenv")
		python.InitPythonVirtualEnv()
	case "typescript":
		log.Print("Init nodejs with typescript empty project")
		typescript.InitTypescriptProject()
	default:
		log.Fatal("Error: Lang not supported, Nestor only support values: [python, typescript]")
		os.Exit(1)
	}
}
