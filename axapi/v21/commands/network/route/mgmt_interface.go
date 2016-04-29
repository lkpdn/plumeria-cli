package route

import (
	"fmt"
	"os"
	"strings"

	"github.com/jessevdk/go-flags"

	. "github.com/lkpdn/plumeria-cli/axapi/util"
	. "github.com/lkpdn/plumeria-cli/axapi/v21/auth"
	. "github.com/lkpdn/plumeria-cli/axapi/v21/util"
)

type MgmtInterface struct{}

func (m *MgmtInterface) Help() string {
	return `
$ ... network.mgmt_interface <subcommand>

To see options for each <subcommand>, try
$ ... network.mgmt_interface <subdommand> --help

Available Subcommands:
	get              : "network.mgmt_interface.get method."
	`
}

func (m *MgmtInterface) Synopsis() string {
	return "network.mgmt_interface method."
}

func (m *MgmtInterface) Run(args []string) int {
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
 * AXAPI(2.1) Method: Network.Route.Mgmt_interface.Get
 */

func (m *MgmtInterface) Get(args []string) (string, error) {
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
		"method":     "network.mgmt_interface.get",
		"session_id": opts.SessionId,
	})
	req := CreateRequest()
	req.QueryString = item
	res, err := Dispatch(req)
	return res, err
}
