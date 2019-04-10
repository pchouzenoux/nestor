package main

import (
	"fmt"
	"os"
)

import (
	"nestor/python"
	"nestor/typescript"
)

func initProject(lang string) {
	switch lang {
	case "python":
		fmt.Println("Init python virtualenv")
		python.InitPythonVirtualEnv()
	case "typescript":
		typescript.InitTypescriptProject()
	default:
		fmt.Println("Error: Lang not supporter, Nestor only support values: [python, typescript]")
		os.Exit(1)
	}
}

func main() {
	// output := flag.String("ouput", "./", "Output directory. Default: ./")
	// flag.Parse()

	switch os.Args[1] {
	case "init":
		initProject(os.Args[2])

	default:
		fmt.Println("Print HELP") // TODO:
	}
}
