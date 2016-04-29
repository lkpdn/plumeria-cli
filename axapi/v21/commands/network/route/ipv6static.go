package route

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/jessevdk/go-flags"

	inObj "github.com/lkpdn/plumeria-cli/axapi/v21/objects/input/network"

	. "github.com/lkpdn/plumeria-cli/axapi/util"
	. "github.com/lkpdn/plumeria-cli/axapi/v21/auth"
	. "github.com/lkpdn/plumeria-cli/axapi/v21/util"
)

type Ipv6static struct{}

func (i *Ipv6static) Help() string {
	return `
$ ... network.route.ipv6static <subcommand>

To see options for each <subcommand>, try
$ ... network.route.ipv6static <subdommand> --help

Available Subcommands:
	getAll              : "network.route.ipv6static.getAll method."
	search              : "network.route.ipv6static.search method."
	`
}

func (i *Ipv6static) Synopsis() string {
	return "network.route.ipv6static method."
}

func (i *Ipv6static) Run(args []string) int {
	if len(args) == 0 {
		fmt.Println(i.Help())
		os.Exit(0)
	}
	var res string
	switch args[0] {
	case "getAll":
		res, _ = i.GetAll(args)
	case "search":
		res, _ = i.Search(args)
	}
	fmt.Println(res)
	return 0
}

/**
 * AXAPI(2.1) Method: Network.Route.Ipv6static.GetAll
 */

func (i *Ipv6static) GetAll(args []string) (string, error) {
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
		"method":     "network.route.ipv6static.getAll",
		"session_id": opts.SessionId,
	})
	req := CreateRequest()
	req.QueryString = item
	res, err := Dispatch(req)
	return res, err
}

/**
 * AXAPI(2.1) Method: Network.Route.Ipv6static.Search
 */

func (i *Ipv6static) Search(args []string) (string, error) {
	// options
	var opts struct {
		GlobalFlags
		Route inObj.Route
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
	params := map[string]map[string]string{"route": {}}
	v := reflect.ValueOf(opts.Route)
	for i := 0; i < v.NumField(); i++ {
		str, ok := v.Field(i).Interface().(string)
		if !ok {
			fmt.Println("invalid param value")
			os.Exit(1)
		}
		if field := v.Type().Field(i).Tag.Get("param"); field != "" && str != "" {
			params["route"][field] = str
		}
	}
	jsonParam, err := json.Marshal(params)
	if err != nil {
		fmt.Println("invalid param value")
		fmt.Println(err.Error())
	}

	// validate
	if len(params["route"]) == 0 {
		fmt.Println(i.Help())
		os.Exit(1)
	}

	// dispatch
	item := CreateUrlValues(map[string]string{
		"method":     "network.route.ipv6static.search",
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
