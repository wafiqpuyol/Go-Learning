package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTP_Server struct {
	Address string `yaml:"address" env-required:"true"`
}

type Config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true" env-default:"production"`
	Storage     string `yaml:"storage" env-required:"true"`
	HTTP_Server `yaml:"http_server" env-required:"true"`
}

func MustLoad() *Config {
	var configPath string

	// reading env variable
	configPath = os.Getenv("CONFIG_PATH")

	// reading flag
	if configPath == "" {
		flags := flag.String("config_path", "", "path to config file")
		flag.Parse()
		configPath = *flags
		if configPath == "" {
			log.Fatal("config_path is required")
		}
	}

	// reading config file
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config_path does not exist: %s", configPath)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("Can't read the config file: %s", err.Error())
	}

	return &cfg
}
