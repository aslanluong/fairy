package inits

import (
	"fmt"
	"os"

	"github.com/aslanluong/fairy/internal/core"
)

func InitConfig(location string, parser core.YAMLConfigParser) (config *core.Config) {
	cfgFile, err := os.Open(location)
	if err != nil {
		fmt.Println("error opening config file,", err)
	}

	config, err = parser.Decode(cfgFile)
	if err != nil {
		fmt.Println("error decoding config file,", err)
	}

	fmt.Println("config file loaded")
	return
}
