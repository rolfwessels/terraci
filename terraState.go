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
		Package: BuildFromFolder(path.Dir(folder), file, []string{}),
	}
	return terraState
}

var files []os.FileInfo

func BuildPackagesFromFolder(folder string) []Package {
	return BuildPackagesFromFolderWithConfig(folder, []string{})
}

func BuildPackagesFromFolderWithConfig(folder string, configs []string) []Package {
	files, err := ioutil.ReadDir(folder)
	if err != nil {
		log.Fatal(err)
	}

	var packages []Package
	From(files).SelectT(func(f os.FileInfo) Package {
		return BuildFromFolder(folder, f, configs)
	}).WhereT(func(p Package) bool {
		return p.Name != "modules" && (len(p.TfFiles) > 0 || len(p.Packages) > 0)
	}).ToSlice(&packages)

	return packages
}

func BuildFromFolder(basePath string, file os.FileInfo, configs []string) Package {
	fullPath := path.Join(basePath, file.Name())

	p := Package{
		Name:          file.Name(),
		Path:          fullPath,
		TfFiles:       []string{},
		TfVars:        []string{},
		ConfigOptions: []string{},
	}
	if file.IsDir() {
		files, err := ioutil.ReadDir(fullPath)
		if err != nil {
			log.Fatal(err)
		}

		for _, file := range files {
			if strings.HasSuffix(file.Name(), FileTerraform) {
				p.TfFiles = append(p.TfFiles, file.Name())
			}
			if strings.HasSuffix(file.Name(), FileTfVariables) {
				p.TfVars = append(p.TfVars, file.Name())
			}
		}
		if len(p.TfVars) > 1 {
			for _, value := range p.TfVars {
				p.ConfigOptions = append(p.ConfigOptions, strings.Replace(value, FileTfVariables, "", 1))
			}
		}
		if len(p.TfFiles) > 0 {
			p.States = ReadCurrentStates(fullPath, configs)
		}

		newConfigs := BuildConfigs(configs, p.ConfigOptions)

		for _, file := range files {

			if file.IsDir() {
				p.Packages = BuildPackagesFromFolderWithConfig(fullPath, newConfigs)
			}
		}
	}

	return p
}

func BuildConfigs(configs []string, options []string) []string {
	if len(configs) == 0 {
		result := []string{}
		for _, new := range options {
			result = append(result, new)
		}
		return result
	}
	if len(options) == 0 {
		return configs
	}
	result := []string{}
	for _, old := range configs {
		for _, new := range options {
			result = append(result, old+"_"+new)
		}
	}
	return result
}

func ReadCurrentStates(basePath string, configs []string) map[string]PackageState {
	m := make(map[string]PackageState)
	cnfs := configs
	if len(cnfs) == 0 {
		cnfs = []string{"global"}
	}

	for _, conf := range cnfs {
		cstate := PackageState{
			State:       StateUnknown,
			Additions:   0,
			Changes:     0,
			Destroys:    0,
			LastUpdated: -1,
			LogContents: []string{""},
		}
		m[conf] = cstate
		fullPath := path.Join(basePath, ".tci/"+conf+FileTciState)
		_, err := os.Stat(fullPath)

		if err == nil {

			// gonfig.GetConf(fullPath, &cstate)
			Dump(cstate, "File read")
			//log.Fatal(err)
		}

	}

	return m
}

func FindFirstPackage(searchIn Package, name string) (Package, bool) {

	if searchIn.Name == name {
		return searchIn, true
	}

	for _, pack := range searchIn.Packages {

		found, isFound := FindFirstPackage(pack, name)
		if isFound {
			return found, isFound
		}
	}
	return Package{}, false
}

type TerraState struct {
	Path    string  `json:"path"`
	Package Package `json:"package"`
}

type Package struct {
	Name          string                  `json:"name,omitempty"`
	Path          string                  `json:"path,omitempty"`
	ConfigOptions []string                `json:"configOptions,omitempty"`
	TfFiles       []string                `json:"tfFiles,omitempty"`
	TfVars        []string                `json:"tfVars,omitempty"`
	States        map[string]PackageState `json:"states,omitempty"`
	Packages      []Package               `json:"packages"`
}

type PackageState struct {
	State       States   `json:"state"`
	Additions   int      `json:"additions,omitempty"`
	Changes     int      `json:"changes,omitempty"`
	Destroys    int      `json:"destroys,omitempty"`
	LastUpdated int64    `json:"lastUpdated,omitempty"`
	LogContents []string `json:"logContents,omitempty"`
}

// States contains all states that a Package can be in
type States int

var (
	// StateUnknown is when the state of the terraform file is Unknown [0]
	StateUnknown States
	// StateScheduled is when the state of the terraform file is Scheduled [1]
	StateScheduled States = 1
	// StatePlanning is when the state of the terraform file is Planning [2]
	StatePlanning States = 2
	// StateApplying is when the state of the terraform file is Applying [3]
	StateApplying States = 3
	// StateDestroying is when the state of the terraform file is Destroying [4]
	StateDestroying States = 4
	// StateApplied is when the state of the terraform file is Applied [5]
	StateApplied States = 5
	// StateFailed is when the state of the terraform file is Failed [6]
	StateFailed States = 6
)

var (
	FileTfVariables string = ".tfvars"
	FileTerraform   string = ".tf"
	FileTciState    string = ".tcistate"
)
