package acl

import (
	"fmt"
	"os"
	"strings"

	"github.com/jessevdk/go-flags"

	. "github.com/lkpdn/plumeria-cli/axapi/util"
	. "github.com/lkpdn/plumeria-cli/axapi/v21/auth"
	. "github.com/lkpdn/plumeria-cli/axapi/v21/util"
)

type Std struct{}

func (a *Std) Help() string {
	return `
$ ... network.acl.std <subcommand>

To see options for each <subcommand>, try
$ ... network.acl.std <subdommand> --help

Available Subcommands:
	getAll              : "network.acl.std.getAll method."
	search              : "network.acl.std.search method."
	`
}

func (a *Std) Synopsis() string {
	return "network.acl.std method."
}

func (a *Std) Run(args []string) int {
	if len(args) == 0 {
		fmt.Println(a.Help())
		os.Exit(0)
	}
	var res string
	switch args[0] {
	case "getAll":
		res, _ = a.GetAll(args)
	case "search":
		res, _ = a.Search(args)
	}
	fmt.Println(res)
	return 0
}

/**
 * AXAPI(2.1) Method: Network.Acl.Std.GetAll
 */

func (a *Std) GetAll(args []string) (string, error) {
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
		"method":     "network.acl.std.getAll",
		"session_id": opts.SessionId,
	})
	req := CreateRequest()
	req.QueryString = item
	res, err := Dispatch(req)
	return res, err
}

/**
 * AXAPI(2.1) Method: Network.Acl.Std.Search
 */

func (a *Std) Search(args []string) (string, error) {
	// options
	var opts struct {
		GlobalFlags
		Id   string `long:"id" description:"id *mutually exclusive"`
		Name string `long:"name" description:"name *mutually exclusive"`
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
	var parameterKey string
	var parameterVal string
	if (opts.Id == "" && opts.Name == "") ||
		(opts.Id != "" && opts.Name != "") {
		fmt.Println(a.Help())
		os.Exit(1)
	}
	if opts.Id != "" {
		parameterKey = "id"
		parameterVal = opts.Id
	}
	if opts.Name != "" {
		parameterKey = "name"
		parameterVal = opts.Name
	}

	// dispatch
	item := CreateUrlValues(map[string]string{
		"method":     "network.acl.std.search",
		"session_id": opts.SessionId,
		parameterKey: parameterVal,
	})
	req := CreateRequest()
	req.QueryString = item
	res, err := Dispatch(req)
	return res, err
}
