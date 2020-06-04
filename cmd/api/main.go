package main

import (
	"fmt"

	"github.com/hack-fan/config"

	"github.com/masschaos/queenbee/models"
)

func main() {
	// load config from env and docker secret
	var c = new(Config)
	config.MustLoad(c)

	models.Init(c.DB.User, c.DB.Password, c.DB.Host, c.DB.Port, c.DB.Name)

	e := router()
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", c.ServerPort)))
}
