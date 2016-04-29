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

type Time struct{}

func (t *Time) Help() string {
	return `
$ ... system.time <subcommand>

To see options for each <subcommand>, try
$ ... system.time <subdommand> --help

Available Subcommands:
	get             : "This method enables you to get the configured system time, date, and timezone."
	`
}

func (t *Time) Synopsis() string {
	return "system.time method."
}

func (t *Time) Run(args []string) int {
	if len(args) == 0 {
		fmt.Println(t.Help())
		os.Exit(0)
	}
	var res string
	switch args[0] {
	case "get":
		res, _ = t.Get(args)
	}
	fmt.Println(res)
	return 0
}

/**
 * AXAPI(2.1) Method: System.Time.Get
 */

func (t *Time) Get(args []string) (string, error) {
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
		"method":     "system.time.get",
		"session_id": opts.SessionId,
	})
	req := CreateRequest()
	req.QueryString = item
	res, err := Dispatch(req)
	return res, err
}
