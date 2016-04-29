package template

import (
	"fmt"
	"os"
	"strings"

	"github.com/jessevdk/go-flags"

	. "github.com/lkpdn/plumeria-cli/axapi/util"
	. "github.com/lkpdn/plumeria-cli/axapi/v21/auth"
	. "github.com/lkpdn/plumeria-cli/axapi/v21/util"
)

type DstIpPersistence struct{}

func (d *DstIpPersistence) Help() string {
	return `
$ ... slb.template.dst_ip_persistence <subcommand>

To see options for each <subcommand>, try
$ ... slb.template.dst_ip_persistence <subdommand> --help

Available Subcommands:
	getAll             : "slb.template.dst_ip_persistence.getAll method."
	search             : "slb.template.dst_ip_persistence.search method."
	`
}

func (d *DstIpPersistence) Synopsis() string {
	return "slb.template.dst_ip_persistence method."
}

func (d *DstIpPersistence) Run(args []string) int {
	if len(args) == 0 {
		fmt.Println(d.Help())
		os.Exit(0)
	}
	var res string
	switch args[0] {
	case "getAll":
		res, _ = d.GetAll(args)
	case "search":
		res, _ = d.Search(args)
	}
	fmt.Println(res)
	return 0
}

/**
 * AXAPI(2.1) Method: Slb.Template.Dst_ip_persistence.GetAll
 */

func (d *DstIpPersistence) GetAll(args []string) (string, error) {
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
		"method":     "slb.template.dst_ip_persistence.getAll",
		"session_id": opts.SessionId,
	})
	req := CreateRequest()
	req.QueryString = item
	res, err := Dispatch(req)
	return res, err
}

/**
 * AXAPI(2.1) Method: Slb.Template.Dst_ip_persistence.Search
 */

func (d *DstIpPersistence) Search(args []string) (string, error) {
	// options
	var opts struct {
		GlobalFlags
		Name string `long:"name" description:"name of slb.template.dst_ip_persistence"`
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
		fmt.Println(d.Help())
		os.Exit(1)
	}

	// dispatch
	item := CreateUrlValues(map[string]string{
		"method":     "slb.template.dst_ip_persistence.search",
		"session_id": opts.SessionId,
		"name":       opts.Name,
	})
	req := CreateRequest()
	req.QueryString = item
	res, err := Dispatch(req)
	return res, err
}
