package cpu

import (
	"fmt"
	"os"
	"strings"

	"github.com/jessevdk/go-flags"

	. "github.com/lkpdn/plumeria-cli/axapi/util"
	. "github.com/lkpdn/plumeria-cli/axapi/v21/auth"
	. "github.com/lkpdn/plumeria-cli/axapi/v21/util"
)

type CurrentUsage struct{}

func (c *CurrentUsage) Help() string {
	return `
$ ... system.device_info.cpu.current_usage <subcommand>

To see options for each <subcommand>, try
$ ... system.device_info.cpu.current_usage <subdommand> --help

Available Subcommands:
	get             : "system.device_info.cpu.current_usage.get method."
	`
}

func (c *CurrentUsage) Synopsis() string {
	return "system.time method."
}

func (c *CurrentUsage) Run(args []string) int {
	if len(args) == 0 {
		fmt.Println(c.Help())
		os.Exit(0)
	}
	var res string
	switch args[0] {
	case "get":
		res, _ = c.Get(args)
	}
	fmt.Println(res)
	return 0
}

/**
 * AXAPI(2.1) Method: System.Device_info.Cpu.Current_usage.Get
 */

func (c *CurrentUsage) Get(args []string) (string, error) {
	// options
	var opts GlobalFlags

	// parse
	old_args := os.Args
	os.Args[0] = strings.Join(os.Args[:(len(os.Args)-len(args)+1)], " ")
	_, err := flags.Parse(&opts)
	if err != nil {
		os.Exit(1)
	}
	os.Args = old_args
	if opts.SessionId == "" {
		opts.SessionId = IssueSessionId()
	}

	// dispatch
	item := CreateUrlValues(map[string]string{
		"method":     "system.device_info.cpu.current_usage.get",
		"session_id": opts.SessionId,
	})
	req := CreateRequest()
	req.QueryString = item
	res, err := Dispatch(req)
	return res, err
}
