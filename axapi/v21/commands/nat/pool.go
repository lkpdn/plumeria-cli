package nat

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jessevdk/go-flags"

	. "github.com/lkpdn/plumeria-cli/axapi/util"
	. "github.com/lkpdn/plumeria-cli/axapi/v21/auth"
	. "github.com/lkpdn/plumeria-cli/axapi/v21/util"
)

type Pool struct{}

func (p *Pool) Help() string {
	return `
$ ... nat.pool <subcommand>

To see options for each <subcommand>, try
$ ... nat.pool <subdommand> --help

Available Subcommands:
	fetchAllStatistics : "nat.pool.fetchAllStatistics method."
	fetchStatistics    : "nat.pool.fetchStatistics method."
	getAll             : "nat.pool.getAll method."
	search             : "nat.pool.search method."
	`
}

func (p *Pool) Synopsis() string {
	return "nat.pool method."
}

func (p *Pool) Run(args []string) int {
	if len(args) == 0 {
		fmt.Println(p.Help())
		os.Exit(0)
	}
	var res string
	switch args[0] {
	case "fetchAllStatistics":
		res, _ = p.FetchAllStatistics(args)
	case "fetchStatistics":
		res, _ = p.FetchStatistics(args)
	case "getAll":
		res, _ = p.GetAll(args)
	case "search":
		res, _ = p.Search(args)
	}
	fmt.Println(res)
	return 0
}

/**
 * AXAPI(2.1) Method: Nat.Pool.FetchAllStatistics
 */

func (p *Pool) FetchAllStatistics(args []string) (string, error) {
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
		"method":     "nat.pool.fetchAllStatistics",
		"session_id": opts.SessionId,
	})
	req := CreateRequest()
	req.QueryString = item
	res, err := Dispatch(req)
	return res, err
}

/**
 * AXAPI(2.1) Method: Nat.Pool.FetchStatistics
 */

func (p *Pool) FetchStatistics(args []string) (string, error) {
	// options
	var opts struct {
		GlobalFlags
		Name string `long:"name" description:"name"`
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
	if opts.Name == "" {
		fmt.Println(p.Help())
		os.Exit(1)
	}

	// dispatch
	item := CreateUrlValues(map[string]string{
		"method":     "nat.pool.fetchStatistics",
		"session_id": opts.SessionId,
		"name":       opts.Name,
	})
	req := CreateRequest()
	req.QueryString = item
	res, err := Dispatch(req)
	return res, err
}

/**
 * AXAPI(2.1) Method: Nat.Pool.GetAll
 */

func (p *Pool) GetAll(args []string) (string, error) {
	// options
	var opts struct {
		GlobalFlags
		IpVersion uint64 `short:"i" long:"ip-version" description:"ip version. (1:IPv4, 2:Ipv6, 3:Ipv4+Ipv6. Scope:1-3, Default:3)"`
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
	if opts.IpVersion == 0 {
		fmt.Println(p.Help())
		os.Exit(1)
	}

	// dispatch
	item := CreateUrlValues(map[string]string{
		"method":     "nat.pool.getAll",
		"session_id": opts.SessionId,
		"ip_version": strconv.FormatUint(opts.IpVersion, 10),
	})
	req := CreateRequest()
	req.QueryString = item
	res, err := Dispatch(req)
	return res, err
}

/**
 * AXAPI(2.1) Method: Nat.Pool.Search
 */

func (p *Pool) Search(args []string) (string, error) {
	// options
	var opts struct {
		GlobalFlags
		Name string `long:"name" description:"name"`
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
	if opts.Name == "" {
		fmt.Println(p.Help())
		os.Exit(1)
	}

	// dispatch
	item := CreateUrlValues(map[string]string{
		"method":     "nat.pool.search",
		"session_id": opts.SessionId,
		"name":       opts.Name,
	})
	req := CreateRequest()
	req.QueryString = item
	res, err := Dispatch(req)
	return res, err
}
