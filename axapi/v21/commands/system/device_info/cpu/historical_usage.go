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

type HistoricalUsage struct{}

func (h *HistoricalUsage) Help() string {
	return `
$ ... system.device_info.cpu.historical_usage <subcommand>

To see options for each <subcommand>, try
$ ... system.device_info.cpu.historical_usage <subdommand> --help

Available Subcommands:
	get             : "system.device_info.cpu.historical_usage.get method."
	`
}

func (h *HistoricalUsage) Synopsis() string {
	return "system.time method."
}

func (h *HistoricalUsage) Run(args []string) int {
	if len(args) == 0 {
		fmt.Println(h.Help())
		os.Exit(0)
	}
	var res string
	switch args[0] {
	case "get":
		res, _ = h.Get(args)
	}
	fmt.Println(res)
	return 0
}

/**
 * AXAPI(2.1) Method: System.Device_info.Cpu.Historical_usage.Get
 */

func (h *HistoricalUsage) Get(args []string) (string, error) {
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
		"method":     "system.device_info.cpu.historical_usage.get",
		"session_id": opts.SessionId,
	})
	req := CreateRequest()
	req.QueryString = item
	res, err := Dispatch(req)
	return res, err
}
