package nat

import (
	"fmt"
	"os"
	"strings"

	"github.com/jessevdk/go-flags"

	. "github.com/lkpdn/plumeria-cli/axapi/util"
	. "github.com/lkpdn/plumeria-cli/axapi/v21/auth"
	. "github.com/lkpdn/plumeria-cli/axapi/v21/util"
)

type StaticTranslation struct{}

func (s *StaticTranslation) Help() string {
	return `
$ ... nat.static_translation <subcommand>

To see options for each <subcommand>, try
$ ... nat.static_translation <subdommand> --help

Available Subcommands:
	fetchAllStatistics : "nat.static_translation.fetchAllStatistics method."
	fetchStatistics    : "nat.static_translation.fetchStatistics method."
	getAll             : "nat.static_translation.getAll method."
	search             : "nat.static_translation.search method."
	`
}

func (s *StaticTranslation) Synopsis() string {
	return "nat.static_translation method."
}

func (s *StaticTranslation) Run(args []string) int {
	if len(args) == 0 {
		fmt.Println(s.Help())
		os.Exit(0)
	}
	var res string
	switch args[0] {
	case "fetchAllStatistics":
		res, _ = s.FetchAllStatistics(args)
	case "fetchStatistics":
		res, _ = s.FetchStatistics(args)
	case "getAll":
		res, _ = s.GetAll(args)
	case "search":
		res, _ = s.Search(args)
	}
	fmt.Println(res)
	return 0
}

/**
 * AXAPI(2.1) Method: Nat.Static_translation.FetchAllStatistics
 */

func (s *StaticTranslation) FetchAllStatistics(args []string) (string, error) {
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
		"method":     "nat.static_translation.fetchAllStatistics",
		"session_id": opts.SessionId,
	})
	req := CreateRequest()
	req.QueryString = item
	res, err := Dispatch(req)
	return res, err
}

/**
 * AXAPI(2.1) Method: Nat.Static_translation.FetchStatistics
 */

func (s *StaticTranslation) FetchStatistics(args []string) (string, error) {
	// options
	var opts struct {
		GlobalFlags
		SourceIpAddr string `long:"src-ip" description:"source ip address"`
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
	if opts.SourceIpAddr == "" {
		fmt.Println(s.Help())
		os.Exit(1)
	}

	// dispatch
	item := CreateUrlValues(map[string]string{
		"method":         "nat.static_translation.fetchStatistics",
		"session_id":     opts.SessionId,
		"source_ip_addr": opts.SourceIpAddr,
	})
	req := CreateRequest()
	req.QueryString = item
	res, err := Dispatch(req)
	return res, err
}

/**
 * AXAPI(2.1) Method: Nat.Static_translation.GetAll
 */

func (s *StaticTranslation) GetAll(args []string) (string, error) {
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
		"method":     "nat.static_translation.getAll",
		"session_id": opts.SessionId,
	})
	req := CreateRequest()
	req.QueryString = item
	res, err := Dispatch(req)
	return res, err
}

/**
 * AXAPI(2.1) Method: Nat.Static_translation.Search
 */

func (s *StaticTranslation) Search(args []string) (string, error) {
	// options
	var opts struct {
		GlobalFlags
		SourceIpAddr string `long:"src-ip" description:"source ip address"`
		GlobalIpAddr string `long:"global-ip" description:"global ip address"`
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
	if opts.SourceIpAddr == "" {
		fmt.Println(s.Help())
		os.Exit(1)
	}
	if opts.GlobalIpAddr == "" {
		fmt.Println(s.Help())
		os.Exit(1)
	}

	// dispatch
	item := CreateUrlValues(map[string]string{
		"method":         "nat.static_translation.search",
		"session_id":     opts.SessionId,
		"source_ip_addr": opts.SourceIpAddr,
		"global_ip_addr": opts.GlobalIpAddr,
	})
	req := CreateRequest()
	req.QueryString = item
	res, err := Dispatch(req)
	return res, err
}
