package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/franela/goreq"
	pretty "github.com/tonnerre/golang-pretty"
)

func PrettifyJson(jsonString string) (string, error) {
	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(jsonString), &dat); err != nil {
		return jsonString, errors.New("Not a JSON string")
	}
	b, err := json.MarshalIndent(dat, "", "  ")
	if err != nil {
		return jsonString, errors.New("JSON string may be corrupted")
	}
	outDat := append(b, '\n')
	return string(outDat), err
}

func Dispatch(req goreq.Request) (string, error) {
	// --debug to show only request uri (plus data if POSTing)
	for _, arg := range os.Args {
		if arg == "-d" || arg == "--debug" {
			httpreq, _ := req.NewRequest()
			fmt.Printf("%# v", pretty.Formatter(httpreq))
			os.Exit(0)
		}
	}

	// dispatch
	res, err := req.Do()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	body, err := res.Body.ToString()

	// --prettify
	for _, arg := range os.Args {
		if arg == "--prettify" {
			body, _ = PrettifyJson(body)
		}
	}
	return body, err
}
