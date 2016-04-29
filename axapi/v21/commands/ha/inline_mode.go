package ha

import (
	"fmt"
	"os"
	"strings"

	"github.com/jessevdk/go-flags"

	. "github.com/lkpdn/plumeria-cli/axapi/util"
	. "github.com/lkpdn/plumeria-cli/axapi/v21/auth"
	. "github.com/lkpdn/plumeria-cli/axapi/v21/util"
)

type InlineMode struct{}

func (i *InlineMode) Help() string {
	return `
$ ... ha.inline_mode <subcommand>

To see options for each <subcommand>, try
$ ... ha.inline_mode <subdommand> --help

Available Subcommands:
	get                 : "ha.inline_mode.get method."
	`
}

func (i *InlineMode) Synopsis() string {
	return "ha.inline_mode method."
}

func (i *InlineMode) Run(args []string) int {
	if len(args) == 0 {
		fmt.Println(i.Help())
		os.Exit(0)
	}
	var res string
	switch args[0] {
	case "get":
		res, _ = i.Get(args)
	}
	fmt.Println(res)
	return 0
}

/**
 * AXAPI(2.1) Method: Ha.Interface.Get
 */

func (i *InlineMode) Get(args []string) (string, error) {
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
		"method":     "ha.inline_mode.get",
		"session_id": opts.SessionId,
	})
	req := CreateRequest()
	req.QueryString = item
	res, err := Dispatch(req)
	return res, err
}
