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
	value := BuildFromFolder("./resources/sample/", file, []string{})
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
	value := BuildFromFolder("./resources/sample/", file, []string{})
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
	value := BuildFromFolder("./resources/sample/", file, []string{})
	if ok, message := So(value.TfFiles, ShouldContain, "main.tf"); !ok {
		t.Error(message)
	}

	if ok, message := So(len(value.TfFiles), ShouldEqual, 3); !ok {
		t.Error(message)
	}
}

func TestBuildFromFolderGivenTwoConfigOptions(t *testing.T) {
	t.Log("Scan for packages")
	file, err := os.Stat("./resources/sample/global")
	if err != nil {
		log.Fatal(err)
	}
	value := BuildFromFolder("./resources/sample/", file, []string{"dev1", "notdev1"})

	if ok, message := So(len(value.States), ShouldEqual, 2); !ok {
		t.Error(message)
	}

	if ok, message := So(len(value.ConfigOptions), ShouldEqual, 0); !ok {
		t.Error(message)
	}

	if ok, message := So(value.States["dev1"].State, ShouldEqual, StateUnknown); !ok {
		t.Error(message)
	}
	_, exists := value.States["foo"]
	if ok, message := So(exists, ShouldBeFalse); !ok {
		t.Error(message)
	}
}

func TestBuildFromFolderGivenNoOptionsConfigOptions(t *testing.T) {
	t.Log("Scan for packages")
	file, err := os.Stat("./resources/sample/global")
	if err != nil {
		log.Fatal(err)
	}
	value := BuildFromFolder("./resources/sample/", file, []string{})

	t.Log("Package with no params should always have a global state")
	if ok, message := So(len(value.States), ShouldEqual, 1); !ok {
		t.Error(message)
	}
	_, exists := value.States["global"]
	if ok, message := So(exists, ShouldBeTrue); !ok {
		t.Error(message)
	}

}

func TestBuildFromFolderGivenTwoConfigOptionsShouldAddThem(t *testing.T) {
	t.Log("Scan for packages")
	file, err := os.Stat("./resources/sample/region")
	if err != nil {
		log.Fatal(err)
	}
	value := BuildFromFolder("./resources/sample/", file, []string{"dev1", "notdev1"})

	if ok, message := So(len(value.ConfigOptions), ShouldEqual, 2); !ok {
		t.Error(message)
	}

}

func TestBuildFromFolderGivenFolderWithMultipleConfigsSubPackageShouldContainNewConfigs(t *testing.T) {
	t.Log("Scan for packages")
	file, err := os.Stat("./resources/sample/region")
	if err != nil {
		log.Fatal(err)
	}
	globalPackage := BuildFromFolder("./resources/sample/", file, []string{})

	appPackage, isFound := FindFirstPackage(globalPackage, "api")

	if ok, message := So(isFound, ShouldBeTrue); !ok {
		t.Error(message)
	}

	t.Log("There should be 6 packages for the 5 config files")
	if ok, message := So(len(appPackage.States), ShouldEqual, 6); !ok {
		t.Error(message)
	}
}
