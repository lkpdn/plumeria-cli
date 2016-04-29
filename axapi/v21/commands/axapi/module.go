package axapi

import (
	"fmt"
	"os"
	"strings"

	"github.com/jessevdk/go-flags"

	. "github.com/lkpdn/plumeria-cli/axapi/util"
	. "github.com/lkpdn/plumeria-cli/axapi/v21/auth"
	. "github.com/lkpdn/plumeria-cli/axapi/v21/util"
)

type Module struct{}

func (m *Module) Help() string {
	return `
$ ... axapi.module <subcommand>

To see options for each <subcommand>, try
$ ... axapi.module <subdommand> --help

Available Subcommands:
	search             : "axapi.module.search method."
	`
}

func (m *Module) Synopsis() string {
	return "axapi.module method."
}

func (m *Module) Run(args []string) int {
	if len(args) == 0 {
		fmt.Println(m.Help())
		os.Exit(0)
	}
	var res string
	switch args[0] {
	case "search":
		res, _ = m.Search(args)
	}
	fmt.Println(res)
	return 0
}

/**
 * AXAPI(2.1) Method: Axapi.Module.Search
 */

func (m *Module) Search(args []string) (string, error) {
	// options
	var opts struct {
		GlobalFlags
		Name       string `long:"name" description:"module name."`
		ApiVersion string `long:"api-version" description:"aXAPI version. (Default:V2.1)"`
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
	if opts.ApiVersion == "" {
		opts.ApiVersion = "V2.1"
	}
	if opts.Name == "" {
		fmt.Println(m.Help())
		os.Exit(1)
	}

	// dispatch
	item := CreateUrlValues(map[string]string{
		"method":     "axapi.module.search",
		"session_id": opts.SessionId,
		"version":    opts.ApiVersion,
		"name":       opts.Name,
	})
	req := CreateRequest()
	req.QueryString = item
	res, err := Dispatch(req)
	return res, err
}
