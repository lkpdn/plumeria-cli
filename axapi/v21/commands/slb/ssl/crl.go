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

type Crl struct{}

func (c *Crl) Help() string {
	return `
$ ... slb.ssl.crl <subcommand>

To see options for each <subcommand>, try
$ ... slb.ssl.crl <subdommand> --help

Available Subcommands:
	getAll    : "slb.ssl.crl.getAll method."
	`
}

func (c *Crl) Synopsis() string {
	return "slb.ssl.crl method."
}

func (c *Crl) Run(args []string) int {
	if len(args) == 0 {
		fmt.Println(c.Help())
		os.Exit(0)
	}
	var res string
	switch args[0] {
	case "getAll":
		res, _ = c.GetAll(args)
	}
	fmt.Println(res)
	return 0
}

/**
 * AXAPI(2.1) Method: Slb.Ssl.Crl.GetAll
 */

func (c *Crl) GetAll(args []string) (string, error) {
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
		"method":     "slb.ssl.crl.getAll",
		"session_id": opts.SessionId,
	})
	req := CreateRequest()
	req.QueryString = item
	res, err := Dispatch(req)
	return res, err
}
