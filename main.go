package main

import (
	"flag"

	"github.com/atmom/repo/common"
	"github.com/atmom/repo/config"
	"github.com/atmom/repo/server"
)

var (
	configPath = flag.String("config", "", "The repository server configuration path")
)

func main() {
	// parse command flags
	flag.Parse()

	// load configuration from json file.
	c := config.Load(*configPath)

	// init log.
	common.InitLog(c)

	// start repository server.
	server.Start(c)
}
