package config

import (
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	App      *App
	Postgres *Postgres
}

type App struct {
	Port   string
	Paging *Paging
}

type Paging struct {
	MaxItemPerPage int
}

type Postgres struct {
	Host              string
	Port              int
	Database          string
	Username          string
	Password          string
	ConnectionTimeout int
	IsPrintLog        bool
}

func Read() *Config {
	conf := &Config{}

	viper.SetConfigFile("config/config.yaml")
	viper.AutomaticEnv()

	// For env variables, so it would be like MYSQL_HOST
	viper.SetEnvKeyReplacer(strings.NewReplacer(`.`, `_`))

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(conf); err != nil {
		panic(err)
	}

	return conf
}
