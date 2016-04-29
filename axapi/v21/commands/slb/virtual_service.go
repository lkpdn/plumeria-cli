package slb

import (
	"fmt"
	"os"
	"strings"

	"github.com/jessevdk/go-flags"

	. "github.com/lkpdn/plumeria-cli/axapi/util"
	. "github.com/lkpdn/plumeria-cli/axapi/v21/auth"
	. "github.com/lkpdn/plumeria-cli/axapi/v21/util"
)

type VirtualService struct{}

func (v *VirtualService) Help() string {
	return `
$ ... slb.virtual_service <subcommand>

To see options for each <subcommand>, try
$ ... slb.virtual_service <subcommand> --help

Available Subcommands:
	fetchAllStatistics : "slb.virtual_service.fetchAllStatistics method."
	fetchStatistics    : "slb.virtual_service.fetchStatistics method."
	getAll             : "slb.virtual_service.getAll method."
	search             : "slb.virtual_service.search method."
	`
}

func (v *VirtualService) Synopsis() string {
	return "slb.virtual_service method."
}

func (v *VirtualService) Run(args []string) int {
	if len(args) == 0 {
		fmt.Println(v.Help())
		os.Exit(0)
	}
	var res string
	switch args[0] {
	case "fetchAllStatistics":
		res, _ = v.FetchAllStatistics(args)
	case "fetchStatistics":
		res, _ = v.FetchStatistics(args)
	case "getAll":
		res, _ = v.GetAll(args)
	case "search":
		res, _ = v.Search(args)
	}
	fmt.Println(res)
	return 0
}

/**
 * AXAPI(2.1) Method: Slb.VirtualService.FetchAllStatistics
 */

func (v *VirtualService) FetchAllStatistics(args []string) (string, error) {
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
		"method":     "slb.virtual_service.fetchAllStatistics",
		"session_id": opts.SessionId,
	})
	req := CreateRequest()
	req.QueryString = item
	res, err := Dispatch(req)
	return res, err
}

/**
 * AXAPI(2.1) Method: Slb.VirtualService.FetchStatistics
 */

func (v *VirtualService) FetchStatistics(args []string) (string, error) {
	// options
	var opts struct {
		GlobalFlags
		Name string `long:"name" description:"name of virtual service"`
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
		fmt.Println(v.Help())
		os.Exit(1)
	}

	// dispatch
	item := CreateUrlValues(map[string]string{
		"method":     "slb.virtual_service.fetchStatistics",
		"session_id": opts.SessionId,
		"name":       opts.Name,
	})
	req := CreateRequest()
	req.QueryString = item
	res, err := Dispatch(req)
	return res, err
}

/**
 * AXAPI(2.1) Method: Slb.VirtualService.GetAll
 */

func (v *VirtualService) GetAll(args []string) (string, error) {
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
		"method":     "slb.virtual_service.getAll",
		"session_id": opts.SessionId,
	})
	req := CreateRequest()
	req.QueryString = item
	res, err := Dispatch(req)
	return res, err
}

/**
 * AXAPI(2.1) Method: Slb.VirtualService.Search
 */

func (v *VirtualService) Search(args []string) (string, error) {
	// options
	var opts struct {
		GlobalFlags
		Name string `long:"name" description:"name of virtual service"`
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
		fmt.Println(v.Help())
		os.Exit(1)
	}

	// dispatch
	item := CreateUrlValues(map[string]string{
		"method":     "slb.virtual_service.search",
		"session_id": opts.SessionId,
		"name":       opts.Name,
	})
	req := CreateRequest()
	req.QueryString = item
	res, err := Dispatch(req)
	return res, err
}
