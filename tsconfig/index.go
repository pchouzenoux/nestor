package tsconfig

import (
	"encoding/json"
)
import (
	"github.com/Nehorim/nestor/commons"
)

type CompilerOptions struct {
	OutDir                           string   `json:"outDir"`
	Target                           string   `json:"target"`
	Lib                              []string `json:"lib"`
	AllowJs                          bool     `json:"allowJs"`
	SkipLibCheck                     bool     `json:"skipLibCheck"`
	EsModuleInterop                  bool     `json:"esModuleInterop"`
	EmitDecoratorMetadata            bool     `json:"emitDecoratorMetadata"`
	ExperimentalDecorators           bool     `json:"experimentalDecorators"`
	AllowSyntheticDefaultImports     bool     `json:"allowSyntheticDefaultImports"`
	Strict                           bool     `json:"strict"`
	ForceConsistentCasingInFileNames bool     `json:"forceConsistentCasingInFileNames"`
	Module                           string   `json:"module"`
	ModuleResolution                 string   `json:"moduleResolution"`
	ResolveJsonModule                bool     `json:"resolveJsonModule"`
	IsolatedModules                  bool     `json:"isolatedModules"`
	InlineSourceMap                  bool     `json:"inlineSourceMap"`
}

type TSConfig struct {
	CompilerOptions  CompilerOptions   `json:"compilerOptions"`
	Files            []string          `json:"files"`
	Include          []string          `json:"include"`
	Exclude          []string          `json:"exclude"`

}

func UnmarshalJsonFile(filepath string) TSConfig {
	var file TSConfig
	var jsonData, _ = commons.ReadFile(filepath)
	json.Unmarshal([]byte(jsonData), &file)
	return file
}

func MarshalJsonFile(pkg TSConfig, filepath string) {
	result, err := json.MarshalIndent(pkg, "", "  ")
	if err != nil {
		panic(err)
	}

	commons.ToFile(filepath, result)
}
