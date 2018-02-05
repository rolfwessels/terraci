package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	. "github.com/ahmetb/go-linq"
	"github.com/gorilla/mux"
	"github.com/urfave/cli"
)

func startCli() {
	app := cli.NewApp()
	app.Name = "continues terraform"
	app.Usage = "Continues deployment for terraform!"
	app.Commands = []cli.Command{
		{
			Name:    "serve",
			Usage:   "serve dashboard for overview",
			Aliases: []string{"s"},
			Action: func(c *cli.Context) error {
				fmt.Println(fmt.Sprintf("Serving on http://localhost:%d/api/terra/state . Press ctr-c to stop.", configuration.Port))
				router := mux.NewRouter()
				router.HandleFunc("/api/terra/state", GetTerraState).Methods("GET")
				log.Printf("INFO: Start serving on http://localhost:%d", configuration.Port)
				log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", configuration.Port), router))
				return nil
			},
		},
		{
			Name:    "plan",
			Aliases: []string{"p"},
			Usage:   "plan individual packages [plan env reg]",
			Action: func(c *cli.Context) error {
				state := BuildTerraState(configuration.BaseTerraformFolder)
				cmdPlan(state, append(c.Args().Tail(), c.Args().First()))
				return nil
			},
		},
	}
	app.Run(os.Args)
}

func cmdPlan(state TerraState, keys []string) {

	commands, err := BuildCommands(state.Package, keys)
	if err != nil {
		log.Fatal("ERROR: ", err)
	}

	if len(commands) == 0 {
		log.Fatalf("ERROR: Could not find any commands to run for '%s'", strings.Join(keys, ","))
	}

	var names []string
	From(commands).SelectT(func(c TerraformCommand) string { return c.Scripts }).ToSlice(&names)
	fmt.Printf("Found %d commands to run.\n", len(names))
	PrintToStdOut(func(w io.Writer) error {
		return TerraformInit(w)
	})
	for _, n := range commands {
		fmt.Printf("=========== %s ============\n", n.Scripts)
		PrintToStdOut(func(w io.Writer) error {
			return TerraformGet(n, w)
		})
		PrintToStdOut(func(w io.Writer) error {
			return TerraformPlan(n, w)
		})
	}
}

type FuncErrorOut func(w io.Writer) error

// GetTerraState returns the state file for a folder
func GetTerraState(w http.ResponseWriter, r *http.Request) {

	terraState := BuildTerraState(configuration.BaseTerraformFolder)
	json.NewEncoder(w).Encode(terraState)
}
