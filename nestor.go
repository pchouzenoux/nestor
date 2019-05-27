package main

import (
	"fmt"
	"log"
	"os"
)

import (
	"github.com/Nehorim/nestor/dockerhelpers"
	"github.com/Nehorim/nestor/githelpers"
	"github.com/Nehorim/nestor/python"
	"github.com/Nehorim/nestor/typescript"
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
		log.Fatal("Error: Command not supported, Try `nestor help` to get more information")
		os.Exit(1)
	}
}

func docker(cmd string) {
	switch cmd {
	case "clean":
		log.Print("Clean docker images")
		dockerhelpers.Clean()
	default:
		log.Fatal("Error: Command not supported, Try `nestor help` to get more information")
		os.Exit(1)
	}
}

func help() {
	fmt.Println("Nestor, your personal butler!")
	fmt.Println("Version: 0.0.1")
	fmt.Println("\nUSAGE:")
	fmt.Println("    nestor <command> <args>")
	fmt.Println("\nCOMMANDS:")
	fmt.Println("    - help    : Print this help")
	fmt.Println("    - init    : Initialize a local project. Args: `python` or `typescript`")
	fmt.Println("    - git     : Git helpers. Args: `clean`")
	fmt.Println("    - docker  : Docker helpers. Args: `clean`")
}

func main() {
	// output := flag.String("ouput", "./", "Output directory. Default: ./")
	// flag.Parse()

	if len(os.Args) != 2 {
		help()
		return
	}
	switch os.Args[1] {
	case "init":
		initProject(os.Args[2])
	case "git":
		git(os.Args[2])
	case "docker":
		docker(os.Args[2])
	default:
		help()
	}
}
