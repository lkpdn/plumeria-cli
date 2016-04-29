package vrrp_a

import (
	"fmt"
	"os"
	"strings"

	"github.com/jessevdk/go-flags"

	. "github.com/lkpdn/plumeria-cli/axapi/util"
	. "github.com/lkpdn/plumeria-cli/axapi/v21/auth"
	. "github.com/lkpdn/plumeria-cli/axapi/v21/util"
)

type VrrpA struct{}

func (v *VrrpA) Help() string {
	return `
$ ... vrrp_a <subcommand>

To see options for each <subcommand>, try
$ ... vrrp_a <subdommand> --help

Available Subcommands:
	get             : "vrrp_a.get method."
	`
}

func (v *VrrpA) Synopsis() string {
	return "vrrp_a method."
}

func (v *VrrpA) Run(args []string) int {
	if len(args) == 0 {
		fmt.Println(v.Help())
		os.Exit(0)
	}
	var res string
	switch args[0] {
	case "get":
		res, _ = v.Get(args)
	}
	fmt.Println(res)
	return 0
}

/**
 * AXAPI(2.1) Method: Vrrp_a.Get
 */

func (v *VrrpA) Get(args []string) (string, error) {
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
		"method":     "vrrp_a.get",
		"session_id": opts.SessionId,
	})
	req := CreateRequest()
	req.QueryString = item
	res, err := Dispatch(req)
	return res, err
}
