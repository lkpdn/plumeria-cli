package session

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/jessevdk/go-flags"

	inObj "github.com/lkpdn/plumeria-cli/axapi/v21/objects/input/slb"

	. "github.com/lkpdn/plumeria-cli/axapi/util"
	. "github.com/lkpdn/plumeria-cli/axapi/v21/auth"
	. "github.com/lkpdn/plumeria-cli/axapi/v21/util"
)

type All struct{}

func (a *All) Help() string {
	return `
$ ... slb.session.all <subcommand>

To see options for each <subcommand>, try
$ ... slb.session.all <subdommand> --help

Available Subcommands:
	get    : "slb.session.all.get method."
	`
}

func (a *All) Synopsis() string {
	return "slb.session.all method."
}

func (a *All) Run(args []string) int {
	if len(args) == 0 {
		fmt.Println(a.Help())
		os.Exit(0)
	}
	var res string
	switch args[0] {
	case "get":
		res, _ = a.Get(args)
	}
	fmt.Println(res)
	return 0
}

/**
 * AXAPI(2.1) Method: Slb.Session.All.Get
 */

func (a *All) Get(args []string) (string, error) {
	// options
	var opts struct {
		GlobalFlags
		SessionFilter inObj.SessionFilter
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

	// construct POST body
	params := map[string]map[string]string{"session_filter": {}}
	v := reflect.ValueOf(opts.SessionFilter)
	for i := 0; i < v.NumField(); i++ {
		if str, ok := v.Field(i).Interface().(string); ok {
			if field := v.Type().Field(i).Tag.Get("param"); field != "" {
				params["session_filter"][field] = str
			}
		} else {
			fmt.Println("invalid param value")
			os.Exit(1)
		}
	}
	jsonParam, err := json.Marshal(params)
	if err != nil {
		fmt.Println("invalid param value")
		fmt.Println(err.Error())
	}

	// dispatch
	item := CreateUrlValues(map[string]string{
		"method":     "slb.session.all.get",
		"session_id": opts.SessionId,
	})
	req := CreateRequest()
	req.Method = "POST"
	req.ContentType = "application/json"
	req.QueryString = item
	req.Body = string(jsonParam)
	res, err := Dispatch(req)
	return res, err
}
