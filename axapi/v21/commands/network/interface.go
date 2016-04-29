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

type Interface struct{}

func (i *Interface) Help() string {
	return `
$ ... network.interface <subcommand>

To see options for each <subcommand>, try
$ ... network.interface <subdommand> --help

Available Subcommands:
	fetchAllStatistics  : "network.interface.fetchAllStatistics method."
	fetchStatistics     : "network.interface.fetchStatistics method."
	getAll              : "network.interface.getAll method."
	get                 : "network.interface.get method."
	`
}

func (i *Interface) Synopsis() string {
	return "network.interface method."
}

func (i *Interface) Run(args []string) int {
	if len(args) == 0 {
		fmt.Println(i.Help())
		os.Exit(0)
	}
	var res string
	switch args[0] {
	case "fetchAllStatistics":
		res, _ = i.FetchAllStatistics(args)
	case "fetchStatistics":
		res, _ = i.FetchStatistics(args)
	case "getAll":
		res, _ = i.GetAll(args)
	case "get":
		res, _ = i.Get(args)
	}
	fmt.Println(res)
	return 0
}

/**
 * AXAPI(2.1) Method: Network.Interface.FetchAllStatistics
 */

func (i *Interface) FetchAllStatistics(args []string) (string, error) {
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
		"method":     "network.interface.fetchAllStatistics",
		"session_id": opts.SessionId,
	})
	req := CreateRequest()
	req.QueryString = item
	res, err := Dispatch(req)
	return res, err
}

/**
 * AXAPI(2.1) Method: Network.Interface.FetchStatistics
 */

func (i *Interface) FetchStatistics(args []string) (string, error) {
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
		fmt.Println(i.Help())
		os.Exit(1)
	}

	// dispatch
	item := CreateUrlValues(map[string]string{
		"method":     "network.interface.fetchStatistics",
		"session_id": opts.SessionId,
		"port_num":   opts.PortNum,
	})
	req := CreateRequest()
	req.QueryString = item
	res, err := Dispatch(req)
	return res, err
}

/**
 * AXAPI(2.1) Method: Network.Interface.GetAll
 */

func (i *Interface) GetAll(args []string) (string, error) {
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
		"method":     "network.interface.getAll",
		"session_id": opts.SessionId,
	})
	req := CreateRequest()
	req.QueryString = item
	res, err := Dispatch(req)
	return res, err
}

/**
 * AXAPI(2.1) Method: Network.Interface.Get
 */

func (i *Interface) Get(args []string) (string, error) {
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
		fmt.Println(i.Help())
		os.Exit(1)
	}

	// dispatch
	item := CreateUrlValues(map[string]string{
		"method":     "network.interface.get",
		"session_id": opts.SessionId,
		"port_num":   opts.PortNum,
	})
	req := CreateRequest()
	req.QueryString = item
	res, err := Dispatch(req)
	return res, err
}
