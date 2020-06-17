package main

import (
	"fmt"

	"github.com/hack-fan/config"
)

func main() {
	// load config from env and docker secret
	var c = new(Config)
	config.MustLoad(c)

	e := router()
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", c.ServerPort)))
}
