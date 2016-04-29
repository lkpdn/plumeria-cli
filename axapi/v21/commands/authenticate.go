package commands

import (
	"fmt"
	. "github.com/lkpdn/plumeria-cli/axapi/util"
	. "github.com/lkpdn/plumeria-cli/axapi/v21/util"
	"github.com/lkpdn/plumeria-cli/config"
)

type Authenticate struct{}

func (a *Authenticate) Help() string {
	return `
... authenticate

No available subcommands. Just return new session id.
`
}

func (a *Authenticate) Synopsis() string {
	return "aXAPI methods for starting and closing aXAPI sessions."
}

func (a *Authenticate) Run(args []string) int {
	res, _ := a.RunTask(args)
	fmt.Println(res)
	return 0
}

func (a *Authenticate) RunTask(args []string) (string, error) {
	username := config.ConfigData.AxDevice.Username
	password := config.ConfigData.AxDevice.Password
	item := CreateUrlValues(map[string]string{
		"method":   "authenticate",
		"username": username,
		"password": password,
	})
	req := CreateRequest()
	req.QueryString = item
	res, err := Dispatch(req)
	return res, err
}
