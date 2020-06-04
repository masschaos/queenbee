package main

// Config can read from env and secret
type Config struct {
	ServerPort string `default:"80"`

	DB struct {
		Name     string `default:"queen"`
		Host     string `required:"true"`
		Port     string `default:"3306"`
		User     string `required:"true"`
		Password string `required:"true"`
	}
}
