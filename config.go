package goui

import (
	"fmt"
	"github.com/fipress/go-rj"
)

var configFile = "config/dev.rj"

type config struct {
	*rj.Node
}

var Config = initConfig()

func initConfig() *config {
	cfg := new(config)

	var err error
	cfg.Node, err = rj.Load(configFile)

	if err != nil {
		fmt.Println("Open config file failed, filename:", configFile, ", error:", err.Error())
	}

	return cfg
}
