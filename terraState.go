package main

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"

	. "github.com/ahmetb/go-linq"
)

func BuildTerraState(folder string) TerraState {
	log.Printf("INFO: Reading from folder %s", folder)
	file, err := os.Stat(folder)
	if err != nil {
		log.Fatal(err)
	}
	terraState := TerraState{
		Path:    folder,
		Package: BuildFromFolder(path.Dir(folder), file),
	}
	return terraState
}

var files []os.FileInfo

// func IsValidBuildFolder(fullPath string) bool {
// 	if file.IsDir() {
// 		files, err := ioutil.ReadDir(file.Name())
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		for _, file := range files {
// 			fmt.Println(file.Name())
// 		}
// 	}
// 	return false
// }

func BuildPackagesFromFolder(folder string) []Package {
	files, err := ioutil.ReadDir(folder)
	if err != nil {
		log.Fatal(err)
	}

	var packages []Package
	From(files).SelectT(func(f os.FileInfo) Package {
		return BuildFromFolder(folder, f)
	}).WhereT(func(p Package) bool {

		return p.Name != "modules" && (len(p.TfFiles) > 0 || len(p.Packages) > 0)
	}).ToSlice(&packages)
	return packages
}

func BuildFromFolder(basePath string, file os.FileInfo) Package {
	fullPath := path.Join(basePath, file.Name())

	p := Package{
		Name:    file.Name(),
		Path:    fullPath,
		TfFiles: []string{},
		TfVars:  []string{},
	}
	if file.IsDir() {
		files, err := ioutil.ReadDir(fullPath)
		if err != nil {
			log.Fatal(err)
		}
		for _, file := range files {
			if strings.HasSuffix(file.Name(), ".tf") {
				p.TfFiles = append(p.TfFiles, file.Name())
			}
			if strings.HasSuffix(file.Name(), ".tfvars") {
				p.TfVars = append(p.TfVars, file.Name())
			}
			if file.IsDir() {
				p.Packages = BuildPackagesFromFolder(fullPath)
			}
		}
	}

	return p
}

type TerraState struct {
	Path    string  `json:"path"`
	Package Package `json:"package"`
}

type Package struct {
	Name     string    `json:"name,omitempty"`
	Path     string    `json:"path,omitempty"`
	TfFiles  []string  `json:"tfFiles,omitempty"`
	TfVars   []string  `json:"tfVars,omitempty"`
	Packages []Package `json:"packages"`
}
