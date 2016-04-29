package vlan

import (
	"fmt"
	"os"
	"strings"

	"github.com/jessevdk/go-flags"

	. "github.com/lkpdn/plumeria-cli/axapi/util"
	. "github.com/lkpdn/plumeria-cli/axapi/v21/auth"
	. "github.com/lkpdn/plumeria-cli/axapi/v21/util"
)

type Mac struct{}

func (m *Mac) Help() string {
	return `
$ ... network.vlan.mac <subcommand>

To see options for each <subcommand>, try
$ ... network.vlan.mac <subdommand> --help

Available Subcommands:
	get              : "network.vlan.mac.get method. get Mac_age_time. (Scope: 10 - 600, Default: 300)"
	`
}

func (m *Mac) Synopsis() string {
	return "network.vlan.mac method."
}

func (m *Mac) Run(args []string) int {
	if len(args) == 0 {
		fmt.Println(m.Help())
		os.Exit(0)
	}
	var res string
	switch args[0] {
	case "get":
		res, _ = m.Get(args)
	}
	fmt.Println(res)
	return 0
}

/**
 * AXAPI(2.1) Method: Network.Vlan.Mac.Get
 */

func (m *Mac) Get(args []string) (string, error) {
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
		"method":     "network.vlan.mac.get",
		"session_id": opts.SessionId,
	})
	req := CreateRequest()
	req.QueryString = item
	res, err := Dispatch(req)
	return res, err
}
