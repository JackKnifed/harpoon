package main

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type tomlConfig struct {
	Port       int    `toml:"port"`
	Addr       string `toml:"addr"`
	Events     map[string]event
	Tunnel     bool
	TunnelName string
}

type event struct {
	Cmd  string `toml:"cmd"`
	Args string `toml:"args"`
}

// loadConfig loads the config.toml file into a tomlConfig struct
func loadConfig() tomlConfig {
	var config tomlConfig
	var err error
	if _, err = toml.DecodeFile(configFile, &config); err != nil {
		fmt.Println(err)
		panic("nope")
	}
	return config
}
