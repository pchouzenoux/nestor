package gitignore

import (
	"log"
	"nestor/commons"
)

var FileContent []byte = []byte(
	`#System Files
.DS_Store

_# Logs
*.log

# IDE - VSCode
.vscode/*
!.vscode/settings.json
!.vscode/tasks.json
!.vscode/launch.json
!.vscode/extensions.json

# misc
.env
`)

var FileName = ".gitignore"

func GetFile() {
	commons.ToFile(FileName, FileContent)
	log.Println("Create file", FileName)
}
