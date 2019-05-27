package npmpack

import (
	"encoding/json"
)
import (
	"nestor/commons"
)

type PackageRepository struct {
	RepoType string `json:"type"`
	Url      string `json:"url"`
}
type PackageBugs struct {
	Url string `json:"url"`
}
type PackageScripts struct {
	Test      string `json:"test"`
	Start     string `json:"start"`
	Build     string `json:"build"`
	Buildlive string `json:"build:live"`
}
type Package struct {
	Name        string            `json:"name"`
	Version     string            `json:"version"`
	Description string            `json:"description"`
	Main        string            `json:"main"`
	Author      string            `json:"author"`
	License     string            `json:"license"`
	Homepage    string            `json:"homepage"`
	Keywords    []string          `json:"keywords"`
	Scripts     PackageScripts    `json:"scripts"`
	Repository  PackageRepository `json:"repository"`
}

func UnmarshalJsonFile(filepath string) Package {
	var pkg Package
	var jsonData, _ = commons.ReadFile(filepath)
	json.Unmarshal([]byte(jsonData), &pkg)
	return pkg
}

func MarshalJsonFile(pkg Package, filepath string) {
	result, err := json.MarshalIndent(pkg, "", "  ")
	if err != nil {
		panic(err)
	}

	commons.ToFile(filepath, result)
}
