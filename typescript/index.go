package typescript

import (
	"strings"
)

import (
	"github.com/Nehorim/nestor/commons"
	"github.com/Nehorim/nestor/npmpack"
	"github.com/Nehorim/nestor/tsconfig"
)

func getTypescriptEnvFile() (string, []byte) {
	var filecontent []byte = []byte(
		`echo nvm use
`)
	return ".env", filecontent
}

func getNvmrcFile() (string, []byte) {
	stdout, _ := commons.ExecShellCmd("node", "-v")
	return ".nvmrc", []byte(stdout)
}

func appendToGitIgnore() (string, []byte) {
	return ".gitignore", []byte(
		`# Dependency directories
node_modules/
`)
}

func setUpNpm() {
	commons.ExecShellCmd("npm", "init -y")

	var pkg = npmpack.UnmarshalJsonFile("package.json")
	pkg.Scripts.Start = `npm run build:live`
	pkg.Scripts.Test = `jest --coverage --colors`
	pkg.Scripts.TestWatch = `jest --watch`
	pkg.Scripts.Build = `tsc -p .`
	pkg.Scripts.Buildlive = `nodemon --watch src/**/*.ts --exec 'npx ts-node' src/index.ts`
	pkg.Scripts.Lint = `tslint -p .`
	pkg.Scripts.LintFix = `tslint --fix -t verbose -p .`
	npmpack.MarshalJsonFile(pkg, "package.json")

	// Install dependencies
	npmInstallDependencies := []string {
		"install",
		"--save",
		"typescript",
		"@types/node",
		"reflect-metadata",
		"source-map-support",
	}
	commons.ExecShellCmd("npm", strings.Join(npmInstallDependencies, " "))

	npmInstallDevDependencies := []string {
		"install",
		"--save-dev",
		"@types/jest",
		"jest",
		"nodemon",
		"prettier",
		"ts-jest",
		"ts-node",
		"tslint",
		"tslint-config-prettier",
		"tslint-eslint-rules",
	}
	commons.ExecShellCmd("npm", strings.Join(npmInstallDevDependencies, " "))
}

func setUpTypescriptConfig() {
	var tsconfigFile = new(tsconfig.TSConfig)
	var compilerOptions = new(tsconfig.CompilerOptions)
	compilerOptions.OutDir = "./dist"
	compilerOptions.Target = "esnext"
	compilerOptions.Lib = []string {"esnext", "esnext.asynciterable"}
	compilerOptions.AllowJs = true
	compilerOptions.SkipLibCheck = true
	compilerOptions.EsModuleInterop = true
	compilerOptions.EmitDecoratorMetadata = true
	compilerOptions.ExperimentalDecorators = true
	compilerOptions.AllowSyntheticDefaultImports = true
	compilerOptions.Strict = false
	compilerOptions.ForceConsistentCasingInFileNames = true
	compilerOptions.Module = "CommonJS"
	compilerOptions.ModuleResolution = "Node"
	compilerOptions.ResolveJsonModule = true
	compilerOptions.IsolatedModules = true
	compilerOptions.InlineSourceMap = true
	
	tsconfigFile.CompilerOptions = *compilerOptions
	tsconfigFile.Files = []string{"./node_modules/@types/node/index.d.ts"}
	tsconfigFile.Include = []string{"src/**/*"}
	tsconfigFile.Exclude = []string{
		"node_modules",
		"src/**/**.test.ts",
		"src/**/**.spec.ts",
		"src/**/__mocks__/**/*",
	}
	tsconfig.MarshalJsonFile(*tsconfigFile, "tsconfig.json")
}

func InitTypescriptProject() {
	commons.ToFile(getNvmrcFile())
	commons.ToFile(getTypescriptEnvFile())

	setUpNpm()

	setUpTypescriptConfig()

	commons.ExecShellCmd("mkdir", "src/")

	commons.AppendToFile(appendToGitIgnore())
}
