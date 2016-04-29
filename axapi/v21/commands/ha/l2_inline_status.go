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

type L2InlineStatus struct{}

func (l *L2InlineStatus) Help() string {
	return `
$ ... ha.l2_inline_status <subcommand>

To see options for each <subcommand>, try
$ ... ha.l2_inline_status <subdommand> --help

Available Subcommands:
	get             : "ha.l2_inline_status.get method."
	`
}

func (l *L2InlineStatus) Synopsis() string {
	return "ha.l2_inline_status method."
}

func (l *L2InlineStatus) Run(args []string) int {
	if len(args) == 0 {
		fmt.Println(l.Help())
		os.Exit(0)
	}
	var res string
	switch args[0] {
	case "get":
		res, _ = l.Get(args)
	}
	fmt.Println(res)
	return 0
}

/**
 * AXAPI(2.1) Method: Nat.L2_inline_status.Get
 */

func (l *L2InlineStatus) Get(args []string) (string, error) {
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
		"method":     "ha.l2_inline_status.get",
		"session_id": opts.SessionId,
	})
	req := CreateRequest()
	req.QueryString = item
	res, err := Dispatch(req)
	return res, err
}
