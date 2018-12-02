package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	. "github.com/ahmetb/go-linq"
)

func TerraformInit(w io.Writer) error {
	if _, err := os.Stat(".terraform/plugins"); err == nil {
		// path/to/whatever exists
		return nil
	} else {
		return execCommand("terraform", []string{"init"}, w)
	}

}

func TerraformGet(command TerraformCommand, w io.Writer) error {

	return execCommand("terraform", []string{"get", command.Scripts}, w)
}

func TerraformPlan(command TerraformCommand, w io.Writer) error {
	args := []string{"plan"}
	//"-var-file=""terraform\default.tfvars"" ")
	var confArgs []string
	From(command.Configs).SelectT(func(s string) string {
		return fmt.Sprintf("-var-file=%s", s)
	}).ToSlice(&confArgs)

	args = append(args, confArgs...)
	stateFile := "-state=" + command.StateFile
	outFile := "-out=" + command.OutputFile
	args = append(args, []string{stateFile, outFile, "-no-color", "-input=false", command.Scripts}...)

	return execCommand("terraform", args, w)
}

func TerraformApply(command TerraformCommand, w io.Writer) error {
	args := []string{"apply"}

	args = append(args, []string{"-no-color", "-input=false", command.OutputFile}...)

	return execCommand("terraform", args, w)
}

func execCommand(executable string, args []string, w io.Writer) error {
	fmt.Println(executable, strings.Join(args, " "))
	cmd := exec.Command(executable, args...)
	cmd.Stdout = w
	cmd.Stderr = w

	if err := cmd.Start(); err != nil {
		return err
	}

	err := cmd.Wait()
	return err
}

func BuildCommands(currentPackage Package, keys []string) ([]TerraformCommand, error) {
	command := []TerraformCommand{}
	sec := BuildPackageSequences(currentPackage, []Package{})
	for _, sec := range sec {
		if contains(keys, sec.Name) {
			cmd := TerraformCommand{}
			selectedConfigs := []string{"ts"}
			for _, pacs := range sec.Packages {
				cmd.Scripts = pacs.Path

				cnfs := pacs.TfVars
				if len(cnfs) > 1 {
					var filterdCnfs []string
					var joinops []string

					From(pacs.TfVars).WhereT(func(s string) bool {
						option := strings.TrimSuffix(s, filepath.Ext(s))
						joinops = append(joinops, option)
						return contains(keys, option)
					}).ToSlice(&filterdCnfs)
					if len(filterdCnfs) == 0 {
						err := fmt.Sprintf("When applying terraform changes to %s you need to select one config [%s].", sec.Name, strings.Join(joinops, ", "))
						return command, errors.New(err)
					}
					if len(filterdCnfs) > 1 {
						err := fmt.Sprintf("When applying terraform changes to %s you need to select only one config [%s].", sec.Name, strings.Join(joinops, ", "))
						return command, errors.New(err)
					}
					selectedConfigs = append(selectedConfigs, strings.Replace(filterdCnfs[0], FileTfVariables, "", 1))
					cnfs = filterdCnfs
				}
				for _, confs := range cnfs {
					fullPath := path.Join(pacs.Path, confs)
					cmd.Configs = append(cmd.Configs, fullPath)

				}
			}

			cmd.StateFile = path.Join(cmd.Scripts, strings.Join(selectedConfigs, "-")+".tfstate")
			cmd.OutputFile = path.Join(cmd.Scripts, strings.Join(selectedConfigs, "-")+".tfout")
			command = append(command, cmd)
		}
	}
	return command, nil
}

func BuildPackageSequences(currentPackage Package, basePackages []Package) []PackageSequences {
	allPackages := []PackageSequences{}
	if len(currentPackage.Packages) > 0 {
		for _, p := range currentPackage.Packages {
			app := append(basePackages, []Package{currentPackage}...)
			vv := BuildPackageSequences(p, app)
			allPackages = append(allPackages, vv...)
		}
	} else {
		allPackages = append(allPackages, PackageSequences{
			Name:     currentPackage.Name,
			Packages: append(basePackages, currentPackage),
		})
	}
	return allPackages
}

// TerraformCommand stores system wide configuration
type TerraformCommand struct {
	Configs    []string
	Scripts    string
	StateFile  string
	OutputFile string
}

// contains packages in sequence
type PackageSequences struct {
	Name     string
	Packages []Package
}
