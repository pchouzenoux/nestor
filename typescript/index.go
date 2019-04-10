package typescript

import (
	"nestor/commons"
)

func getTypescriptEnvFile() (string, []byte) {
	var filecontent []byte = []byte(
		`echo nvm use
`)
	return ".env", filecontent
}

func getNvmrcFile() (string, []byte) {
	return ".nvmrc", []byte(commons.ExecShellCmd("node", "-v"))
}

func InitTypescriptProject() {
	commons.ToFile(getNvmrcFile())
	commons.ToFile(getTypescriptEnvFile())

	commons.ExecShellCmd("npm", "init", "-y")
	commons.ExecShellCmd("npm", "install", "--save-dev", "typescript", "@types/node", "nodemon")

	commons.ExecShellCmd("npx", "tsc", "--init", "--rootDir", "src", "--outDir", "lib", "--esModuleInterop", "--resolveJsonModule", "--lib", "es6,dom", "--module", "commonjs")

	commons.ExecShellCmd("mkdir", "src/")
}
