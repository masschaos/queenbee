package main

import (
	"fmt"

	"github.com/masschaos/queenbee/config"
	"github.com/masschaos/queenbee/models"
	"github.com/masschaos/queenbee/route"
)

func main() {
	c := config.GetConfig()
	models.Init(c.DB.User, c.DB.Password, c.DB.Hostname, c.DB.Port, c.DB.DDName)

	e := route.New()
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", c.ServerPort)))
}