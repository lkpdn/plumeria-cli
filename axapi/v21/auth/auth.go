package auth

import (
	"encoding/json"
	"fmt"
	"os"
	"github.com/lkpdn/plumeria-cli/axapi/v21/commands"
)

type ResponseSessionId struct {
	SessionId string `json:"session_id"`
}

func IssueSessionId() string {
	auth := commands.Authenticate{}

	// ugly workaround
	oldArgs := os.Args
	dummyArgs := make([]string, 0, len(os.Args))
	for _, val := range oldArgs {
		if val != "--debug" && val != "-d" {
			dummyArgs = append(dummyArgs, val)
		}
	}
	os.Args = dummyArgs
	res, err := auth.RunTask([]string{""})
	os.Args = oldArgs

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	var data ResponseSessionId
	err = json.Unmarshal([]byte(res), &data)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return data.SessionId
}
