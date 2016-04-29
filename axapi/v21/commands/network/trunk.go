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

type Trunk struct{}

func (t *Trunk) Help() string {
	return `
$ ... network.trunk <subcommand>

To see options for each <subcommand>, try
$ ... network.trunk <subdommand> --help

Available Subcommands:
	fetchAllStatistics  : "network.trunk.fetchAllStatistics method."
	fetchStatisticsById : "network.trunk.fetchStatisticsById method."
	getAll              : "network.trunk.getAll method."
	search              : "network.trunk.search method."
	`
}

func (t *Trunk) Synopsis() string {
	return "network.trunk method."
}

func (t *Trunk) Run(args []string) int {
	if len(args) == 0 {
		fmt.Println(t.Help())
		os.Exit(0)
	}
	var res string
	switch args[0] {
	case "fetchAllStatistics":
		res, _ = t.FetchAllStatistics(args)
	case "fetchStatisticsById":
		res, _ = t.FetchStatisticsById(args)
	case "getAll":
		res, _ = t.GetAll(args)
	case "search":
		res, _ = t.Search(args)
	}
	fmt.Println(res)
	return 0
}

/**
 * AXAPI(2.1) Method: Network.Trunk.FetchAllStatistics
 */

func (t *Trunk) FetchAllStatistics(args []string) (string, error) {
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
		"method":     "network.trunk.fetchAllStatistics",
		"session_id": opts.SessionId,
	})
	req := CreateRequest()
	req.QueryString = item
	res, err := Dispatch(req)
	return res, err
}

/**
 * AXAPI(2.1) Method: Network.Trunk.FetchStatisticsById
 */

func (t *Trunk) FetchStatisticsById(args []string) (string, error) {
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
		fmt.Println(t.Help())
		os.Exit(1)
	}

	// dispatch
	item := CreateUrlValues(map[string]string{
		"method":     "network.trunk.fetchStatisticsById",
		"session_id": opts.SessionId,
		"id":         opts.Id,
	})
	req := CreateRequest()
	req.QueryString = item
	res, err := Dispatch(req)
	return res, err
}

/**
 * AXAPI(2.1) Method: Network.Trunk.GetAll
 */

func (t *Trunk) GetAll(args []string) (string, error) {
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
		"method":     "network.trunk.getAll",
		"session_id": opts.SessionId,
	})
	req := CreateRequest()
	req.QueryString = item
	res, err := Dispatch(req)
	return res, err
}

/**
 * AXAPI(2.1) Method: Network.Trunk.Search
 */

func (t *Trunk) Search(args []string) (string, error) {
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
		fmt.Println(t.Help())
		os.Exit(1)
	}

	// dispatch
	item := CreateUrlValues(map[string]string{
		"method":     "network.trunk.search",
		"session_id": opts.SessionId,
		"id":         opts.Id,
	})
	req := CreateRequest()
	req.QueryString = item
	res, err := Dispatch(req)
	return res, err
}
