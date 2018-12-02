package main

import (
	"io"
	"io/ioutil"
	"strings"
	"testing"

	. "github.com/ahmetb/go-linq"
	. "github.com/smartystreets/assertions"
)

func TestTerraformExec(t *testing.T) {
	pr, pw := io.Pipe()
	defer pr.Close()
	var err error
	go func() {
		// close the writer, so the reader knows there's no more data
		defer pw.Close()
		err = execCommand("terraform", []string{"help"}, pw)
	}()

	slurp, _ := ioutil.ReadAll(pr)

	if ok, message := So(len(slurp), ShouldBeGreaterThan, 1000); !ok {
		t.Error(message)
	}
	if ok, message := So(err.Error(), ShouldEqual, "exit status 127"); !ok {
		t.Error(message)
	}
}

func TestBuildCommands(t *testing.T) {
	state := BuildTerraState("./resources/sample")
	args := []string{"api", "application", "dev", "eu-west-1"}
	x, _ := BuildCommands(state.Package, args)
	if ok, message := So(len(x), ShouldEqual, 2); !ok {
		t.Error(message)
	}
}

func TestBuildCommandsShouldAddScript(t *testing.T) {
	state := BuildTerraState("./resources/sample")
	args := []string{"global"}
	commands, _ := BuildCommands(state.Package, args)
	command := From(commands).Last().(TerraformCommand)
	if ok, message := So(command.Scripts, ShouldEqual, "resources/sample/global"); !ok {
		t.Error(message)
	}
}

func TestBuildCommandsShouldAddConfigs(t *testing.T) {
	state := BuildTerraState("./resources/sample")
	args := []string{"api", "eu-west-1", "dev"}
	commands, _ := BuildCommands(state.Package, args)
	command := From(commands).Last().(TerraformCommand)

	if ok, message := So(command.Configs, ShouldContain, "resources/sample/default.tfvars"); !ok {
		t.Error(message)
	}
	if ok, message := So(command.Configs, ShouldContain, "resources/sample/region/eu-west-1.tfvars"); !ok {
		t.Error(message)
	}
	if ok, message := So(command.Configs, ShouldContain, "resources/sample/region/env/dev.tfvars"); !ok {
		t.Error(message)
	}
	if ok, message := So(len(command.Configs), ShouldEqual, 3); !ok {
		t.Error(message)
	}
}

func TestBuildCommandsShouldHaveStateFile(t *testing.T) {
	state := BuildTerraState("./resources/sample")
	args := []string{"api", "eu-west-1", "dev"}
	commands, _ := BuildCommands(state.Package, args)
	command := From(commands).Last().(TerraformCommand)

	if ok, message := So(command.StateFile, ShouldEqual, "resources/sample/region/env/app/api/ts-eu-west-1-dev.tfstate"); !ok {
		t.Error(message)
	}

}

func TestBuildCommandsShouldFailIfConfigIsNotSelected(t *testing.T) {
	state := BuildTerraState("./resources/sample")
	args := []string{"api", "eu-west-1"}
	_, err := BuildCommands(state.Package, args)
	errorString := err.Error()
	if ok, message := So(errorString, ShouldEqual, "When applying terraform changes to api you need to select one config [dev, prod, uat]."); !ok {
		t.Error(message)
	}
}

func TestBuildPackageSequences(t *testing.T) {
	state := BuildTerraState("./resources/sample")
	seq := BuildPackageSequences(state.Package, []Package{})
	var names []string
	From(seq).SelectT(func(x PackageSequences) string { return x.Name }).ToSlice(&names)

	t.Log("Found the following names:", strings.Join(names, ", "))
	if ok, message := So(len(names), ShouldEqual, 13); !ok {
		t.Error(message)
	}

	if ok, message := So(names, ShouldContain, "api"); !ok {
		t.Error(message)
	}
}

func TestBuildPackageSequencesShouldContainParents(t *testing.T) {
	state := BuildTerraState("./resources/sample")
	seq := BuildPackageSequences(state.Package, []Package{})

	sample := From(seq).FirstWithT(func(x PackageSequences) bool { return x.Name == "api" })
	original, ok := sample.(PackageSequences)
	if ok {
		if ok, message := So(len(original.Packages), ShouldEqual, 5); !ok {
			t.Error(message)
		}
	} else {
		t.Error("could not cast sample")
	}
}
