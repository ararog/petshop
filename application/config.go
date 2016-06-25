package application

import (
	"fmt"
	"os"
	"github.com/BurntSushi/toml"
)

type Config struct {
	DB Database `toml:"database"`
	Server Server
}

type Server struct {
	Port int
}

type Database struct {
	Type string
  ConnectionString string `toml:"connection_string"`
}

func Environment() string {
	environment := os.Getenv("PETSHOP_ENV")
	if environment == "" {
		environment = "development"
	}

	return environment
}

func LoadConfig() Config {
	var config Config

	environment := Environment()
	file := fmt.Sprintf("./config/%s.toml", environment)
	if _, ioerr := os.Stat(file); ioerr == nil {
		if _, err := toml.DecodeFile(file, &config); err != nil {
			panic("failed to read configuration")
		}
	} else {
		panic("enviroment configuration does not exists")
	}

	return config
}
