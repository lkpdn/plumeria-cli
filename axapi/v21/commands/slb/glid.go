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

type Glid struct{}

func (g *Glid) Help() string {
	return `
$ ... slb.glid <subcommand>

To see options for each <subcommand>, try
$ ... slb.glid <subdommand> --help

Available Subcommands:
	getAll             : "slb.glid.getAll method."
	search             : "slb.glid.search method."
	`
}

func (g *Glid) Synopsis() string {
	return "slb.glid method."
}

func (g *Glid) Run(args []string) int {
	if len(args) == 0 {
		fmt.Println(g.Help())
		os.Exit(0)
	}
	var res string
	switch args[0] {
	case "getAll":
		res, _ = g.GetAll(args)
	case "search":
		res, _ = g.Search(args)
	}
	fmt.Println(res)
	return 0
}

/**
 * AXAPI(2.1) Method: Slb.Glid.GetAll
 */

func (g *Glid) GetAll(args []string) (string, error) {
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
		"method":     "slb.glid.getAll",
		"session_id": opts.SessionId,
	})
	req := CreateRequest()
	req.QueryString = item
	res, err := Dispatch(req)
	return res, err
}

/**
 * AXAPI(2.1) Method: Slb.Glid.Search
 */

func (g *Glid) Search(args []string) (string, error) {
	// options
	var opts struct {
		GlobalFlags
		Id string `long:"id" description:"glid id. Scope: 1 - 1023, Default: 0."`
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
		fmt.Println(g.Help())
		os.Exit(1)
	}

	// dispatch
	item := CreateUrlValues(map[string]string{
		"method":     "slb.glid.search",
		"session_id": opts.SessionId,
		"id":         opts.Id,
	})
	req := CreateRequest()
	req.QueryString = item
	res, err := Dispatch(req)
	return res, err
}
