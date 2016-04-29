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

type AclBind struct{}

func (a *AclBind) Help() string {
	return `
$ ... nat.acl_bind <subcommand>

To see options for each <subcommand>, try
$ ... nat.acl_bind <subdommand> --help

Available Subcommands:
	getAll             : "nat.acl_bind.getAll method."
	`
}

func (a *AclBind) Synopsis() string {
	return "nat.acl_bind method."
}

func (a *AclBind) Run(args []string) int {
	if len(args) == 0 {
		fmt.Println(a.Help())
		os.Exit(0)
	}
	var res string
	switch args[0] {
	case "getAll":
		res, _ = a.GetAll(args)
	}
	fmt.Println(res)
	return 0
}

/**
 * AXAPI(2.1) Method: Nat.Acl_bind.GetAll
 */

func (a *AclBind) GetAll(args []string) (string, error) {
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
		"method":     "nat.acl_bind.getAll",
		"session_id": opts.SessionId,
	})
	req := CreateRequest()
	req.QueryString = item
	res, err := Dispatch(req)
	return res, err
}
