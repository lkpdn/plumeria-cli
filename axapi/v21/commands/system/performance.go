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

type Performance struct{}

func (p *Performance) Help() string {
	return `
$ ... system.performance <subcommand>

To see options for each <subcommand>, try
$ ... system.performance <subdommand> --help

Available Subcommands:
	get : system.performance.get method.
	`
}

func (p *Performance) Synopsis() string {
	return "system.performance method."
}

func (p *Performance) Run(args []string) int {
	if len(args) == 0 {
		fmt.Println(p.Help())
		os.Exit(0)
	}
	var res string
	switch args[0] {
	case "get":
		res, _ = p.Get(args)
	}
	fmt.Println(res)
	return 0
}

/**
 * AXAPI(2.1) Method: AXAPI Method: System.Performance.Get
 */

func (p *Performance) Get(args []string) (string, error) {
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
		"method":     "system.performance.get",
		"session_id": opts.SessionId,
	})
	req := CreateRequest()
	req.QueryString = item
	res, err := Dispatch(req)
	return res, err
}
