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

type Cert struct{}

func (c *Cert) Help() string {
	return `
$ ... slb.ssl.cert <subcommand>

To see options for each <subcommand>, try
$ ... slb.ssl.cert <subdommand> --help

Available Subcommands:
	getAll    : "slb.ssl.cert.getAll method."
	`
}

func (c *Cert) Synopsis() string {
	return "slb.ssl.cert method."
}

func (c *Cert) Run(args []string) int {
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
 * AXAPI(2.1) Method: Slb.Ssl.Cert.GetAll
 */

func (c *Cert) GetAll(args []string) (string, error) {
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
		"method":     "slb.ssl.cert.getAll",
		"session_id": opts.SessionId,
	})
	req := CreateRequest()
	req.QueryString = item
	res, err := Dispatch(req)
	return res, err
}
