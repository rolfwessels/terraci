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
		return execCommand("terraform.exe", []string{"init"}, w)
	}

}

func TerraformGet(command TerraformCommand, w io.Writer) error {
	return execCommand("terraform.exe", []string{"get", command.Scripts}, w)
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
	args = append(args, []string{stateFile, "-no-color", "-input=false", command.Scripts}...)

	return execCommand("terraform.exe", args, w)
}

func execCommand(executable string, args []string, w io.Writer) error {
	fmt.Println(executable, strings.Join(args, " "))
	cmd := exec.Command(executable, args...)
	cmd.Stdout = w
	cmd.Stderr = w

	if err := cmd.Start(); err != nil {
		return err
	}

	if err := cmd.Wait(); err != nil {
		return err
	}
	return nil
}

func BuildCommands(currentPackage Package, keys []string) ([]TerraformCommand, error) {
	command := []TerraformCommand{}
	sec := BuildPackageSequences(currentPackage, []Package{})
	for _, sec := range sec {
		if contains(keys, sec.Name) {
			cmd := TerraformCommand{}
			keys := []string{}
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

					cnfs = filterdCnfs
				}
				for _, confs := range cnfs {
					fullPath := path.Join(pacs.Path, confs)
					cmd.Configs = append(cmd.Configs, fullPath)
					//keys = append(keys, confs)
				}
			}

			cmd.StateFile = path.Join(cmd.Scripts, "ts"+strings.Join(keys, "-")+".tfstate")
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
	Configs   []string
	Scripts   string
	StateFile string
}

// contains packages in sequence
type PackageSequences struct {
	Name     string
	Packages []Package
}
