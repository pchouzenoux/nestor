package main

import (
	"log"
	"os"
)

import (
	"nestor/dockerhelpers"
	"nestor/githelpers"
	"nestor/python"
	"nestor/typescript"
)

func initProject(lang string) {
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

func git(cmd string) {
	switch cmd {
	case "clean":
		log.Print("Clean git local repository")
		githelpers.Clean()
	default:
		log.Fatal("Error: Command not supported, Try `nestor --help` to get more information")
		os.Exit(1)
	}
}

func docker(cmd string) {
	switch cmd {
	case "clean":
		log.Print("Clean docker images")
		dockerhelpers.Clean()
	default:
		log.Fatal("Error: Command not supported, Try `nestor --help` to get more information")
		os.Exit(1)
	}
}

func main() {
	// output := flag.String("ouput", "./", "Output directory. Default: ./")
	// flag.Parse()

	switch os.Args[1] {
	case "init":
		initProject(os.Args[2])
	case "git":
		git(os.Args[2])
	case "docker":
		docker(os.Args[2])

	default:
		log.Println("Print HELP") // TODO:
	}
}
