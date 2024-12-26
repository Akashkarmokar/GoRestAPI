package config

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HttpServer struct {
	Addr string `yaml:"address" env-required:"true"`
}

type Config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true" env-default:"production"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HttpServer  `yaml:"http_server"`
}

func MustLoad() *Config {
	var configPath string
	configPath = os.Getenv("CONFIG_PATH") // Get From Env File

	if configPath == "" { // Check also from flags or arguments
		flags := flag.String("config", "", "Path to the configuration file")
		flag.Parse()

		configPath = *flags

		fmt.Println("ConfigPath: ", configPath)

		if configPath == "" {
			log.Fatal("Config path is not exist")
		}
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config file does not exist: %s", configPath)
	}

	var cfg Config
	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("Can not read config file: %s", err.Error())
	}

	return &cfg
}
