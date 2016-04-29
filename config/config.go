package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type AxDevice struct {
	Ip       string
	Username string
	Password string
}

type Config struct {
	AxDevice AxDevice `json:"ax_device"`
}

var ConfigData Config

func ConfigInit(filepath ...string) {
	if filepath == nil || filepath[0] != "" {
		file, err := ioutil.ReadFile(filepath[0])
		if err != nil {
			fmt.Printf("Could not read config JSON file: %v\n", filepath)
			fmt.Printf("error: %v\n", err)
			os.Exit(1)
		}
		err = json.Unmarshal(file, &ConfigData)
		if err != nil {
			fmt.Println("JSON file might not conform to required format.")
			fmt.Printf("error: %v\n", err)
			os.Exit(1)
		}
	} else {
		ConfigData = Config{
			AxDevice: AxDevice{
				Ip:       "10.10.10.10",
				Username: "admin",
				Password: "admin",
			},
		}
	}
}
