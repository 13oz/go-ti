package main

import (
	"./config"
)

func main() {
	config.ReadConfigFile("config.toml")
}
