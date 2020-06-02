package config

import (
	xconfig "github.com/hack-fan/config"
)

type Config struct {
	ServerPort string `default:"80" env:"SERVER_PORT"`

	DB struct {
		DDName   string `env:"MYSQL_DB_NAME"`
		Hostname string `env:"MYSQL_HOSTNAME"`
		Port     string `default:"3306" env:"MYSQL_PORT"`
		User     string `env:"MYSQL_USER"`
		Password string `env:"MYSQL_PASSWORD"`
	}
}

func GetConfig() *Config {
	c := &Config{}
	err := xconfig.Load(c)
	if err != nil {
		panic(err)
	}

	return c
}
