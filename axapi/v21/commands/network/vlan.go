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

type Vlan struct{}

func (v *Vlan) Help() string {
	return `
$ ... network.vlan <subcommand>

To see options for each <subcommand>, try
$ ... network.vlan <subdommand> --help

Available Subcommands:
	fetchAllStatistics  : "network.vlan.fetchAllStatistics method."
	fetchStatisticsById : "network.vlan.fetchStatisticsById method."
	getAll              : "network.vlan.getAll method."
	search              : "network.vlan.search method."
	`
}

func (v *Vlan) Synopsis() string {
	return "network.vlan method."
}

func (v *Vlan) Run(args []string) int {
	if len(args) == 0 {
		fmt.Println(v.Help())
		os.Exit(0)
	}
	var res string
	switch args[0] {
	case "fetchAllStatistics":
		res, _ = v.FetchAllStatistics(args)
	case "fetchStatisticsById":
		res, _ = v.FetchStatisticsById(args)
	case "getAll":
		res, _ = v.GetAll(args)
	case "search":
		res, _ = v.Search(args)
	}
	fmt.Println(res)
	return 0
}

/**
 * AXAPI(2.1) Method: Network.Vlan.FetchAllStatistics
 */

func (v *Vlan) FetchAllStatistics(args []string) (string, error) {
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
		"method":     "network.vlan.fetchAllStatistics",
		"session_id": opts.SessionId,
	})
	req := CreateRequest()
	req.QueryString = item
	res, err := Dispatch(req)
	return res, err
}

/**
 * AXAPI(2.1) Method: Network.Vlan.FetchStatisticsById
 */

func (v *Vlan) FetchStatisticsById(args []string) (string, error) {
	// options
	var opts struct {
		GlobalFlags
		VlanId string `long:"vlan-id" description:"vlan id"`
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
	if opts.VlanId == "" {
		fmt.Println(v.Help())
		os.Exit(1)
	}

	// dispatch
	item := CreateUrlValues(map[string]string{
		"method":     "network.vlan.fetchStatisticsById",
		"session_id": opts.SessionId,
		"vlan_id":    opts.VlanId,
	})
	req := CreateRequest()
	req.QueryString = item
	res, err := Dispatch(req)
	return res, err
}

/**
 * AXAPI(2.1) Method: Network.Vlan.GetAll
 */

func (v *Vlan) GetAll(args []string) (string, error) {
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
		"method":     "network.vlan.getAll",
		"session_id": opts.SessionId,
	})
	req := CreateRequest()
	req.QueryString = item
	res, err := Dispatch(req)
	return res, err
}

/**
 * AXAPI(2.1) Method: Network.Vlan.Search
 */

func (v *Vlan) Search(args []string) (string, error) {
	// options
	var opts struct {
		GlobalFlags
		Id string `long:"id" description:"id"`
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
	if opts.Id == "" {
		fmt.Println(v.Help())
		os.Exit(1)
	}

	// dispatch
	item := CreateUrlValues(map[string]string{
		"method":     "network.vlan.search",
		"session_id": opts.SessionId,
		"id":         opts.Id,
	})
	req := CreateRequest()
	req.QueryString = item
	res, err := Dispatch(req)
	return res, err
}
