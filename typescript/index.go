package typescript

import (
	"nestor/commons"
	"nestor/npmpack"
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

func appendToGitIgnore() (string, []byte) {
	return ".gitignore", []byte(
		`# Dependency directories
node_modules/
`)
}

func InitTypescriptProject() {
	commons.ToFile(getNvmrcFile())
	commons.ToFile(getTypescriptEnvFile())

	commons.ExecShellCmd("npm", "init -y")

	var pkg = npmpack.UnmarshalJsonFile("package.json")
	pkg.Scripts.Start = `npm run build:live`
	pkg.Scripts.Build = `tsc -p .`
	pkg.Scripts.Buildlive = `nodemon --watch src/**/*.ts --exec 'npx ts-node' src/index.ts`
	npmpack.MarshalJsonFile(pkg, "package.json")

	commons.ExecShellCmd("npm", "install --save-dev typescript @types/node nodemon")

	commons.ExecShellCmd("npx", "tsc --init --rootDir src --outDir lib --esModuleInterop --resolveJsonModule --lib es6,dom --module commonjs")

	commons.ExecShellCmd("mkdir", "src/")

	commons.AppendToFile(appendToGitIgnore())
}
