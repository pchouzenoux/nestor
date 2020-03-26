package typescript

import (
	"strings"
)

import (
	"github.com/Nehorim/nestor/commons"
)

func getTypescriptEnvFile() (string, []byte) {
	var filecontent []byte = []byte(
		`nvm use
`)
	return ".env", filecontent
}

func getESLintIgnoreFile() (string, []byte) {
	var filecontent []byte = []byte(
		`dist/
`)
	return ".eslintignore", filecontent
}

func getESLintConfigFile() (string, []byte) {
	var filecontent []byte = []byte(
		`module.exports = {
			parser: '@typescript-eslint/parser', // Specifies the ESLint parser
			parserOptions: {
			  tsconfigRootDir: __dirname,
			  extends: './tsconfig.json',
			},
			env: {
			  jest: true,
			  node: true,
			},
			plugins: ['prettier'],
			extends: [
			  'airbnb-base',
			  'eslint:recommended',
			  'plugin:import/errors',
			  'plugin:import/warnings',
			  'plugin:@typescript-eslint/eslint-recommended',
			  'plugin:@typescript-eslint/recommended',
			  'plugin:import/typescript',
			  'prettier',
			  'prettier/@typescript-eslint',
			  'plugin:prettier/recommended',
			],
			rules: {
			  // Place to specify ESLint rules. Can be used to overwrite rules specified from the extended configs
			  // e.g. "@typescript-eslint/explicit-function-return-type": "off",
			  'sort-imports': [
				'warn',
				{
				  ignoreCase: true,
				  ignoreDeclarationSort: true,
				  ignoreMemberSort: false,
				},
			  ],
			  '@typescript-eslint/no-explicit-any': 'off',
			  '@typescript-eslint/camelcase': 'off',
			  // Override Airbnb base config
			  'import/first': 'off',
			  'no-useless-constructor': 'off',
			  'class-methods-use-this': 'off',
			  'import/prefer-default-export': 'off',
			  'import/extensions': ['error', 'never'],
			  'max-classes-per-file': 'off',
			  'no-restricted-syntax': [
				'error',
				{
				  selector: 'ForInStatement',
				  message:
					'for..in loops iterate over the entire prototype chain, which is virtually never what you want. Use Object.{keys,values,entries}, and iterate over the resulting array.',
				},
				{
				  selector: 'LabeledStatement',
				  message:
					'Labels are a form of GOTO; using them makes code confusing and hard to maintain and understand.',
				},
				{
				  selector: 'WithStatement',
				  message:
					"'with' is disallowed in strict mode because it makes code impossible to predict and optimize.",
				},
			  ],
			},
		  };`)
	return ".eslintrc.js", filecontent
}

func getPrettierConfigFile() (string, []byte) {
	var filecontent []byte = []byte(
		`{
			"requireConfig": true,
			"singleQuote": true,
			"trailingComma": "all"
		  }
		  `)
	return ".prettierrc", filecontent
}

func getJestConfigFile() (string, []byte) {
	var filecontent []byte = []byte(
		`module.exports = {
			roots: ["<rootDir>/src"],
			setupFiles: ["reflect-metadata"],
			transform: {
			  "^.+\\.ts?$": "ts-jest"
			}
		  };		  
		  `)
	return "jest.config.js", filecontent
}

func getNvmrcFile() (string, []byte) {
	stdout, _ := commons.ExecShellCmd("node", "-v")
	return ".nvmrc", []byte(stdout)
}

func appendToGitIgnore() (string, []byte) {
	return ".gitignore", []byte(
		`# Dependency directories
node_modules/
coverage/
`)
}

func setUpNpm() {
	commons.ExecShellCmd("npm", "init -y")

	var pkg = UnmarshalNPMFile("package.json")
	pkg.Scripts.Start = `yarn run build:live`
	pkg.Scripts.Test = `jest --coverage --colors`
	pkg.Scripts.Build = `rm -rf dist/ && tsc -p ./tsconfig.build.json`
	pkg.Scripts.Buildlive = `tsnd --watch 'src/**/*.ts,config/**/*.*' src/index.ts`
	pkg.Scripts.Lint = `eslint --ext .ts,.js .`
	MarshalNPMFile(pkg, "package.json")

	// Install dependencies
	npmInstallDependencies := []string {
		"add",
		"typescript",
		"reflect-metadata",
		"source-map-support",
	}
	commons.ExecShellCmd("yarn", strings.Join(npmInstallDependencies, " "))

	npmInstallDevDependencies := []string {
		"add",
		"--dev",
		"@types/jest",
		"@types/node",
		"jest",
		"ts-node-dev",
		"prettier",
		"ts-jest",
		"ts-node",
		"eslint",
		"eslint-config-airbnb-base",
		"eslint-config-prettier",
		"eslint-plugin-import",
		"eslint-plugin-prettier",
	}
	commons.ExecShellCmd("yarn", strings.Join(npmInstallDevDependencies, " "))
}

func setUpTypescriptConfig() {
	var tsconfigFile = new(TSConfig)
	var compilerOptions = new(TSCompilerOptions)
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
	MarshalTSConfigFile(*tsconfigFile, "tsconfig.json")
}

func InitTypescriptProject() {
	commons.ToFile(getNvmrcFile())
	commons.ToFile(getTypescriptEnvFile())
	commons.ToFile(getESLintConfigFile())
	commons.ToFile(getESLintIgnoreFile())
	commons.ToFile(getPrettierConfigFile())
	commons.ToFile(getJestConfigFile())

	setUpNpm()

	setUpTypescriptConfig()

	commons.ExecShellCmd("mkdir", "src/")

	commons.AppendToFile(appendToGitIgnore())
}
