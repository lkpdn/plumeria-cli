package network

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jessevdk/go-flags"

	. "github.com/lkpdn/plumeria-cli/axapi/util"
	. "github.com/lkpdn/plumeria-cli/axapi/v21/auth"
	. "github.com/lkpdn/plumeria-cli/axapi/v21/util"
)

type Arp struct{}

func (a *Arp) Help() string {
	return `
$ ... network.arp <subcommand>

To see options for each <subcommand>, try
$ ... network.arp <subdommand> --help

Available Subcommands:
	fetchAllStatistics  : "network.arp.fetchAllStatistics method."
	fetchStatisticsByIp : "network.arp.fetchStatisticsByIp method."
	getAll              : "network.arp.getAll method."
	search              : "network.arp.search method."
	`
}

func (a *Arp) Synopsis() string {
	return "network.arp method."
}

func (a *Arp) Run(args []string) int {
	if len(args) == 0 {
		fmt.Println(a.Help())
		os.Exit(0)
	}
	var res string
	switch args[0] {
	case "fetchAllStatistics":
		res, _ = a.FetchAllStatistics(args)
	case "fetchStatisticsByIp":
		res, _ = a.FetchStatisticsByIp(args)
	case "getAll":
		res, _ = a.GetAll(args)
	case "search":
		res, _ = a.Search(args)
	}
	fmt.Println(res)
	return 0
}

/**
 * AXAPI(2.1) Method: Network.Arp.FetchAllStatistics
 */

func (a *Arp) FetchAllStatistics(args []string) (string, error) {
	// options
	var opts struct {
		GlobalFlags
		IpVersion uint64 `short:"i" long:"ip-version" description:"ip version. 4 or 6"`
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
	if opts.IpVersion == 0 {
		fmt.Println(a.Help())
		os.Exit(1)
	}

	// dispatch
	item := CreateUrlValues(map[string]string{
		"method":     "network.arp.fetchAllStatistics",
		"session_id": opts.SessionId,
		"ip_version": strconv.FormatUint(opts.IpVersion, 10),
	})
	req := CreateRequest()
	req.QueryString = item
	res, err := Dispatch(req)
	return res, err
}

/**
 * AXAPI(2.1) Method: Network.Arp.FetchStatisticsByIp
 */

func (a *Arp) FetchStatisticsByIp(args []string) (string, error) {
	// options
	var opts struct {
		GlobalFlags
		Ipv4Address string `long:"ipv4-address" description:"ip address(ipv4) *mutually exclusive"`
		Ipv6Address string `long:"ipv6-address" description:"ip address(ipv6) *mutually exclusive"`
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
	var ipAddressKey string
	var ipAddressVal string
	if (opts.Ipv4Address == "" && opts.Ipv6Address == "") ||
		(opts.Ipv4Address != "" && opts.Ipv6Address != "") {
		fmt.Println(a.Help())
		os.Exit(1)
	}
	if opts.Ipv4Address != "" {
		ipAddressKey = "ipv4_address"
		ipAddressVal = opts.Ipv4Address
	}
	if opts.Ipv6Address != "" {
		ipAddressKey = "ipv6_address"
		ipAddressVal = opts.Ipv6Address
	}

	// dispatch
	item := CreateUrlValues(map[string]string{
		"method":     "network.arp.fetchStatisticsByIp",
		"session_id": opts.SessionId,
		ipAddressKey: ipAddressVal,
	})
	req := CreateRequest()
	req.QueryString = item
	res, err := Dispatch(req)
	return res, err
}

/**
 * AXAPI(2.1) Method: Network.Arp.GetAll
 */

func (a *Arp) GetAll(args []string) (string, error) {
	// options
	var opts struct {
		GlobalFlags
		IpVersion uint64 `short:"i" long:"ip-version" description:"ip version. 4 or 6"`
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
	if opts.IpVersion == 0 {
		fmt.Println(a.Help())
		os.Exit(1)
	}

	// dispatch
	item := CreateUrlValues(map[string]string{
		"method":     "network.arp.getAll",
		"session_id": opts.SessionId,
		"ip_version": strconv.FormatUint(opts.IpVersion, 10),
	})
	req := CreateRequest()
	req.QueryString = item
	res, err := Dispatch(req)
	return res, err
}

/**
 * AXAPI(2.1) Method: Network.Arp.Search
 */

func (a *Arp) Search(args []string) (string, error) {
	// options
	var opts struct {
		GlobalFlags
		Ipv4Address string `long:"ipv4-address" description:"ip address(ipv4) *mutually exclusive"`
		Ipv6Address string `long:"ipv6-address" description:"ip address(ipv6) *mutually exclusive"`
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
	var ipAddressKey string
	var ipAddressVal string
	if (opts.Ipv4Address == "" && opts.Ipv6Address == "") ||
		(opts.Ipv4Address != "" && opts.Ipv6Address != "") {
		fmt.Println(a.Help())
		os.Exit(1)
	}
	if opts.Ipv4Address != "" {
		ipAddressKey = "ipv4_address"
		ipAddressVal = opts.Ipv4Address
	}
	if opts.Ipv6Address != "" {
		ipAddressKey = "ipv6_address"
		ipAddressVal = opts.Ipv6Address
	}

	// dispatch
	item := CreateUrlValues(map[string]string{
		"method":     "network.arp.search",
		"session_id": opts.SessionId,
		ipAddressKey: ipAddressVal,
	})
	req := CreateRequest()
	req.QueryString = item
	res, err := Dispatch(req)
	return res, err
}
