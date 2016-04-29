package system

import (
	"fmt"
	"os"
	"strings"

	"github.com/jessevdk/go-flags"

	. "github.com/lkpdn/plumeria-cli/axapi/util"
	. "github.com/lkpdn/plumeria-cli/axapi/v21/auth"
	. "github.com/lkpdn/plumeria-cli/axapi/v21/util"
)

type Resource struct{}

func (r *Resource) Help() string {
	return `
$ ... system.resource <subcommand>

To see options for each <subcommand>, try
$ ... system.resource <subdommand> --help

Available Subcommands:
	get             : "system.resource.get method."
	`
}

func (r *Resource) Synopsis() string {
	return "system.resource method."
}

func (r *Resource) Run(args []string) int {
	if len(args) == 0 {
		fmt.Println(r.Help())
		os.Exit(0)
	}
	var res string
	switch args[0] {
	case "get":
		res, _ = r.Get(args)
	}
	fmt.Println(res)
	return 0
}

/**
 * AXAPI(2.1) Method: System.Resource.Get
 */

func (r *Resource) Get(args []string) (string, error) {
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
		"method":     "system.resource.get",
		"session_id": opts.SessionId,
	})
	req := CreateRequest()
	req.QueryString = item
	res, err := Dispatch(req)
	return res, err
}
