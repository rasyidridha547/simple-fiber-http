package config

import (
	"log"
	"sync"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port string `envconfig:"PORT"`
}

var cfg Config
var once sync.Once

func Get() Config {
	once.Do(func() {
		if err := envconfig.Process("", &cfg); err != nil {
			log.Fatalf("Failed to load config: %+v", err.Error())
		}
	})

	return cfg
}
