package main

import (
	"encoding/json"
	"io/ioutil"
)

func getConfig() Config {
	data, _ := ioutil.ReadFile("/usr/local/main/config.json")
	var c Config
	_ = json.Unmarshal(data, &c)
	return c
}
