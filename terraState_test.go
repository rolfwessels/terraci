package main

import (
	"log"
	"os"
	"testing"

	. "github.com/smartystreets/assertions"
)

func TestWhenBuildingStateShouldSetPath(t *testing.T) {
	t.Log("Settings path should add path to state.")
	value := BuildTerraState("./resources/sample")
	if ok, message := So(value.Path, ShouldEqual, "./resources/sample"); !ok {
		t.Error(message)
	}
}

func TestBuildStateWhenCalledWithValidPathShouldReadTerraformFiles(t *testing.T) {
	t.Log("Scan for packages")
	value := BuildTerraState("./resources/sample")
	if ok, message := So(value.Package.Name, ShouldEqual, "sample"); !ok {
		t.Error(message)
	}
}

func TestBuildFromFolderShouldReadName(t *testing.T) {
	t.Log("Scan for packages")
	file, err := os.Stat("./resources/sample/global")
	if err != nil {
		log.Fatal(err)
	}
	value := BuildFromFolder("./resources/sample/", file)
	if ok, message := So(value.Name, ShouldEqual, "global"); !ok {
		t.Error(message)
	}
}

func TestBuildFromFolderShouldSetPath(t *testing.T) {
	t.Log("Scan for packages")
	file, err := os.Stat("./resources/sample/global")
	if err != nil {
		log.Fatal(err)
	}
	value := BuildFromFolder("./resources/sample/", file)
	if ok, message := So(value.Path, ShouldEqual, "resources/sample/global"); !ok {
		t.Error(message)
	}
}

func TestBuildFromFolderShouldSetReadTfFiles(t *testing.T) {
	t.Log("Scan for packages")
	file, err := os.Stat("./resources/sample/global")
	if err != nil {
		log.Fatal(err)
	}
	value := BuildFromFolder("./resources/sample/", file)
	if ok, message := So(value.TfFiles, ShouldContain, "main.tf"); !ok {
		t.Error(message)
	}
	if ok, message := So(len(value.TfFiles), ShouldEqual, 1); !ok {
		t.Error(message)
	}
}
