package global

import (
	"fmt"
	"os"
	"strings"

	"github.com/jessevdk/go-flags"

	. "github.com/lkpdn/plumeria-cli/axapi/util"
	. "github.com/lkpdn/plumeria-cli/axapi/v21/auth"
	. "github.com/lkpdn/plumeria-cli/axapi/v21/util"
)

type LogRateLimiting struct{}

func (l *LogRateLimiting) Help() string {
	return `
$ ... slb.global.log_rate_limiting <subcommand>

To see options for each <subcommand>, try
$ ... slb.global.log_rate_limiting <subdommand> --help

Available Subcommands:
	get              : "slb.global.log_rate_limiting.get method."
	`
}

func (l *LogRateLimiting) Synopsis() string {
	return "slb.global.log_rate_limiting method."
}

func (l *LogRateLimiting) Run(args []string) int {
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
 * AXAPI(2.1) Method: Slb.Global.Log_rate_limiting.Get
 */

func (l *LogRateLimiting) Get(args []string) (string, error) {
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
		"method":     "slb.global.log_rate_limiting.get",
		"session_id": opts.SessionId,
	})
	req := CreateRequest()
	req.QueryString = item
	res, err := Dispatch(req)
	return res, err
}
