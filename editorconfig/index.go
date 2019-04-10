package editorconfig

import (
	"log"
	"nestor/commons"
)

var FileContent []byte = []byte(
	`# Editor configuration, see http://editorconfig.org
root = true

[*]
charset = utf-8
indent_style = space
indent_size = 2
insert_final_newline = true
trim_trailing_whitespace = true

[*.md]
max_line_length = off
trim_trailing_whitespace = false
`)

var FileName = ".editorconfig"

func GetFile() {
	commons.ToFile(FileName, FileContent)
	log.Println("Create file", FileName)
}
