package ssl

import (
	"fmt"
	"os"
	"strings"

	"github.com/jessevdk/go-flags"

	. "github.com/lkpdn/plumeria-cli/axapi/util"
	. "github.com/lkpdn/plumeria-cli/axapi/v21/auth"
	. "github.com/lkpdn/plumeria-cli/axapi/v21/util"
)

type ExpirationEmail struct{}

func (e *ExpirationEmail) Help() string {
	return `
$ ... slb.ssl.expiration_email <subcommand>

To see options for each <subcommand>, try
$ ... slb.ssl.expiration_email <subdommand> --help

Available Subcommands:
	get      : "slb.ssl.expiration_email.get method."
	`
}

func (e *ExpirationEmail) Synopsis() string {
	return "slb.ssl.expiration_email method."
}

func (e *ExpirationEmail) Run(args []string) int {
	if len(args) == 0 {
		fmt.Println(e.Help())
		os.Exit(0)
	}
	var res string
	switch args[0] {
	case "get":
		res, _ = e.Get(args)
	}
	fmt.Println(res)
	return 0
}

/**
 * AXAPI(2.1) Method: Slb.Ssl.Expiration_email.Get
 */

func (e *ExpirationEmail) Get(args []string) (string, error) {
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
		"method":     "slb.ssl.expiration_email.get",
		"session_id": opts.SessionId,
	})
	req := CreateRequest()
	req.QueryString = item
	res, err := Dispatch(req)
	return res, err
}
