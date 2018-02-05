package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/tkanos/gonfig"
)

var (
	configuration Configuration
)

func main() {

	// Setup configuration
	configuration = loadConfigFile()

	// Setup logging to write to file
	f, err := os.OpenFile(configuration.Logfile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	mw := io.MultiWriter(os.Stdout, f)
	log.SetOutput(mw)

	startCli()
}

// Configuration stores system wide configuration
type Configuration struct {
	Port                int
	BaseTerraformFolder string
	Logfile             string
	Terraform           string
}

func loadConfigFile() Configuration {
	conf := Configuration{
		Port:                8000,
		BaseTerraformFolder: "casdcasd",
	}
	gonfig.GetConf("config/config.json", &conf)
	return conf
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func Dump(s interface{}, d string) {
	res2B, _ := json.MarshalIndent(s, ">", "  ")
	fmt.Println(d, string(res2B))
}

func PrintToStdOut(callBack FuncErrorOut) {
	pr, pw := io.Pipe()
	defer pr.Close()
	var err error
	go func() {
		// close the writer, so the reader knows there's no more data
		defer pw.Close()

		err = callBack(pw)
	}()
	scanner := bufio.NewScanner(pr)
	for scanner.Scan() {
		fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	if err != nil {
		log.Fatal(err)
	}
}
