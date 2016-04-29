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

type Interface struct{}

func (i *Interface) Help() string {
	return `
$ ... ha.interface <subcommand>

To see options for each <subcommand>, try
$ ... ha.interface <subdommand> --help

Available Subcommands:
	getAll              : "ha.interface.getAll method."
	get                 : "ha.interface.get method."
	`
}

func (i *Interface) Synopsis() string {
	return "ha.interface method."
}

func (i *Interface) Run(args []string) int {
	if len(args) == 0 {
		fmt.Println(i.Help())
		os.Exit(0)
	}
	var res string
	switch args[0] {
	case "getAll":
		res, _ = i.GetAll(args)
	case "get":
		res, _ = i.Get(args)
	}
	fmt.Println(res)
	return 0
}

/**
 * AXAPI(2.1) Method: Ha.Interface.GetAll
 */

func (i *Interface) GetAll(args []string) (string, error) {
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
		"method":     "ha.interface.getAll",
		"session_id": opts.SessionId,
	})
	req := CreateRequest()
	req.QueryString = item
	res, err := Dispatch(req)
	return res, err
}

/**
 * AXAPI(2.1) Method: Ha.Interface.Get
 */

func (i *Interface) Get(args []string) (string, error) {
	// options
	var opts struct {
		GlobalFlags
		IfNum string `long:"if-num" description:"if number"`
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
	if opts.IfNum == "" {
		fmt.Println(i.Help())
		os.Exit(1)
	}

	// dispatch
	item := CreateUrlValues(map[string]string{
		"method":     "ha.interface.get",
		"session_id": opts.SessionId,
		"if_num":     opts.IfNum,
	})
	req := CreateRequest()
	req.QueryString = item
	res, err := Dispatch(req)
	return res, err
}
