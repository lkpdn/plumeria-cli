package util

import (
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/lkpdn/plumeria-cli/config"

	"github.com/franela/goreq"
	"github.com/jessevdk/go-flags"
)

type GlobalFlags struct {
	SessionId  string `long:"session-id" description:"session id (optional)"`
	Debug      []bool `short:"d" long:"debug" description:"debug"`
	Prettify   []bool `long:"prettify" description:"prettify JSON output"`
	ConfigJson string `long:"config-json" description:"config JSON file path"`
}

func CreateUrlValues(params map[string]string) url.Values {
	item := url.Values{}
	item.Set("format", "json")
	for k, v := range params {
		item.Add(k, v)
	}
	return item
}

func CreateRequest() goreq.Request {
	endpoint := fmt.Sprintf("http://%s/services/rest/V2.1/",
		config.ConfigData.AxDevice.Ip)
	req := goreq.Request{
		Uri:         endpoint,
		Accept:      "application/json",
		ContentType: "application/json",
		Insecure:    true,
	}
	return req
}

func SessionIdParse(args []string) string {
	var opts struct {
		SessionId string `long:"session-id" description:"session id"`
	}
	old_args := os.Args
	os.Args[0] = strings.Join(os.Args[:(len(os.Args)-len(args)+1)], " ")
	_, err := flags.Parse(&opts)
	if err != nil {
		os.Exit(1)
	}
	os.Args = old_args
	return opts.SessionId
}
