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

type Ext struct{}

func (e *Ext) Help() string {
	return `
$ ... network.acl.ext <subcommand>

To see options for each <subcommand>, try
$ ... network.acl.ext <subdommand> --help

Available Subcommands:
	getAll              : "network.acl.ext.getAll method."
	search              : "network.acl.ext.search method."
	`
}

func (e *Ext) Synopsis() string {
	return "network.acl.ext method."
}

func (e *Ext) Run(args []string) int {
	if len(args) == 0 {
		fmt.Println(e.Help())
		os.Exit(0)
	}
	var res string
	switch args[0] {
	case "getAll":
		res, _ = e.GetAll(args)
	case "search":
		res, _ = e.Search(args)
	}
	fmt.Println(res)
	return 0
}

/**
 * AXAPI(2.1) Method: Network.Acl.Ext.GetAll
 */

func (e *Ext) GetAll(args []string) (string, error) {
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
		"method":     "network.acl.ext.getAll",
		"session_id": opts.SessionId,
	})
	req := CreateRequest()
	req.QueryString = item
	res, err := Dispatch(req)
	return res, err
}

/**
 * AXAPI(2.1) Method: Network.Acl.Ext.Search
 */

func (e *Ext) Search(args []string) (string, error) {
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
		fmt.Println(e.Help())
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
		"method":     "network.acl.ext.search",
		"session_id": opts.SessionId,
		parameterKey: parameterVal,
	})
	req := CreateRequest()
	req.QueryString = item
	res, err := Dispatch(req)
	return res, err
}
