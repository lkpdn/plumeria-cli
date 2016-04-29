package template

import (
	"fmt"
	"os"
	"strings"

	"github.com/jessevdk/go-flags"

	. "github.com/lkpdn/plumeria-cli/axapi/util"
	. "github.com/lkpdn/plumeria-cli/axapi/v21/auth"
	. "github.com/lkpdn/plumeria-cli/axapi/v21/util"
)

type Sip struct{}

func (s *Sip) Help() string {
	return `
$ ... slb.template.sip <subcommand>

To see options for each <subcommand>, try
$ ... slb.template.sip <subdommand> --help

Available Subcommands:
	getAll             : "slb.template.sip.getAll method."
	search             : "slb.template.sip.search method."
	`
}

func (s *Sip) Synopsis() string {
	return "slb.template.sip method."
}

func (s *Sip) Run(args []string) int {
	if len(args) == 0 {
		fmt.Println(s.Help())
		os.Exit(0)
	}
	var res string
	switch args[0] {
	case "getAll":
		res, _ = s.GetAll(args)
	case "search":
		res, _ = s.Search(args)
	}
	fmt.Println(res)
	return 0
}

/**
 * AXAPI(2.1) Method: Slb.Template.Sip.GetAll
 */

func (s *Sip) GetAll(args []string) (string, error) {
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
		"method":     "slb.template.sip.getAll",
		"session_id": opts.SessionId,
	})
	req := CreateRequest()
	req.QueryString = item
	res, err := Dispatch(req)
	return res, err
}

/**
 * AXAPI(2.1) Method: Slb.Template.Sip.Search
 */

func (s *Sip) Search(args []string) (string, error) {
	// options
	var opts struct {
		GlobalFlags
		Name string `long:"name" description:"name of slb.template.sip"`
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
	if opts.Name == "" {
		fmt.Println(s.Help())
		os.Exit(1)
	}

	// dispatch
	item := CreateUrlValues(map[string]string{
		"method":     "slb.template.sip.search",
		"session_id": opts.SessionId,
		"name":       opts.Name,
	})
	req := CreateRequest()
	req.QueryString = item
	res, err := Dispatch(req)
	return res, err
}
