package network

import (
	"fmt"
	"os"
	"strings"

	"github.com/jessevdk/go-flags"

	. "github.com/lkpdn/plumeria-cli/axapi/util"
	. "github.com/lkpdn/plumeria-cli/axapi/v21/auth"
	. "github.com/lkpdn/plumeria-cli/axapi/v21/util"
)

type Ve struct{}

func (v *Ve) Help() string {
	return `
$ ... network.ve <subcommand>

To see options for each <subcommand>, try
$ ... network.ve <subdommand> --help

Available Subcommands:
	get    : "network.ve.get method."
	getAll : "network.ve.getAll method."
	`
}

func (v *Ve) Synopsis() string {
	return "network.ve method."
}

func (v *Ve) Run(args []string) int {
	if len(args) == 0 {
		fmt.Println(v.Help())
		os.Exit(0)
	}
	var res string
	switch args[0] {
	case "get":
		res, _ = v.Get(args)
	case "getAll":
		res, _ = v.GetAll(args)
	}
	fmt.Println(res)
	return 0
}

/**
 * AXAPI(2.1) Method: Network.Ve.Get
 */

func (v *Ve) Get(args []string) (string, error) {
	// options
	var opts struct {
		GlobalFlags
		PortNum string `long:"port-num" description:"port number"`
	}

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

	// validate
	if opts.PortNum == "" {
		fmt.Println(v.Help())
		os.Exit(1)
	}

	// dispatch
	item := CreateUrlValues(map[string]string{
		"method":     "network.ve.get",
		"session_id": opts.SessionId,
		"port_num":   opts.PortNum,
	})
	req := CreateRequest()
	req.QueryString = item
	res, err := Dispatch(req)
	return res, err
}

/**
 * AXAPI(2.1) Method: Network.Ve.GetAll
 */

func (v *Ve) GetAll(args []string) (string, error) {
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
		"method":     "network.ve.getAll",
		"session_id": opts.SessionId,
	})
	req := CreateRequest()
	req.QueryString = item
	res, err := Dispatch(req)
	return res, err
}
