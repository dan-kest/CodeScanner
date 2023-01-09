package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	App      *App
	Postgres *Postgres
	RabbitMQ *RabbitMQ
}

type App struct {
	Port   string
	Paging *Paging
	Scan   *Scan
}

type Paging struct {
	MaxItemPerPage int
}

type Scan struct {
	LocalRepoPath string
	WorkerCount   int
	Ignore        string
	FindingRule   []FindingRule
}

type FindingRule struct {
	Type        string
	Match       string
	RuleID      string
	Description string
	Severity    string
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

type RabbitMQ struct {
	Host           string
	Port           int
	Username       string
	Password       string
	PublishTimeout int
	Queue          *Queue
}

type Queue struct {
	Name string
}

func Read() *Config {
	conf := &Config{}

	viper.SetConfigFile("config/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	viper.SetConfigFile("config/rule.yaml")
	if err := viper.MergeInConfig(); err != nil {
		panic(err)
	}

	viper.SetConfigFile("config/secret.yaml")
	if err := viper.MergeInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(conf); err != nil {
		panic(err)
	}

	return conf
}

// For testing, returns empty struct with no nil values.
func ReadDummy() *Config {
	return &Config{
		App: &App{
			Paging: &Paging{},
			Scan: &Scan{
				FindingRule: []FindingRule{},
			},
		},
		Postgres: &Postgres{},
		RabbitMQ: &RabbitMQ{
			Queue: &Queue{},
		},
	}
}
