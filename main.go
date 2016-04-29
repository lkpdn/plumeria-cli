package main

import (
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/lkpdn/plumeria-cli/axapi/v21"
	"github.com/lkpdn/plumeria-cli/config"

	"github.com/mitchellh/cli"
)

const (
	CLI_VERSION = "0.0.1"
)

func main() {
	AXAPI_VERSIONS := map[string]map[string]cli.CommandFactory{
		"v21": v21.Commands,
	}

	command := os.Args[0]
	args := os.Args[1:]

	// --version or -v to just show this cli tool version.
	for _, arg := range args {
		if arg == "-v" || arg == "--version" {
			fmt.Println(CLI_VERSION)
			os.Exit(0)
		}
	}

	// aXAPI version string must be specified.
	api_version := ""
	var command_map map[string]cli.CommandFactory
	if len(args) > 0 {
		for version, cmd_map := range AXAPI_VERSIONS {
			if args[0] == version {
				api_version = version
				command_map = cmd_map
				break
			}
		}
	}
	if api_version == "" {
		fmt.Println("You have to specify api version. try:")
		fmt.Printf("$ %v [API_VERSION] --help\n\n", command)
		fmt.Println("Available API_VERSIONs are:")
		for version, _ := range AXAPI_VERSIONS {
			fmt.Printf("   - %v\n", version)
		}
		os.Exit(0)
	}

	// --config-json to override default ConfigData
	configRe := regexp.MustCompile("--config-json=([^ ]+)")
	var configFilePath string
	for _, arg := range args {
		if group := configRe.FindStringSubmatch(arg); len(group) == 2 {
			configFilePath = group[1]
		}
	}
	config.ConfigInit(configFilePath)

	// run cli
	c := cli.NewCLI(command+" "+api_version, CLI_VERSION)
	c.Args = args[1:]
	c.Commands = command_map
	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}
	os.Exit(exitStatus)
}
