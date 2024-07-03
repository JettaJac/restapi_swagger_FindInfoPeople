package config

import (
	"flag"

	// "fmt"
	// "github.com/caarlos0/env/v6"
	"github.com/ilyakaznacheev/cleanenv"
	// "gopkg.in/yaml.v3"
	// "log"
	"main/internal/storage/postgre"
	"os"
	"time"
)

type Config struct {
	Env          string `yaml:"env" env-default:"local"`
	AuthBase     string `yaml:"authbase" env-default:""`
	NameDataBase string `yaml:"namebase" env-requered:"true"`
	// DatabaseURL  string `yaml:"namebase" env:"DATABASE_HOST" env-default:"localhost" env-required:"true"`
	HTTPServer `yaml:"http_server"`
	Postgres   postgre.Postgres `yaml:"postgres"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:":8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
	// User        string        `yaml:"user" env-required:"true"`
	// Password    string        `yaml:"password" env-required:"true" env:"HTTP_SERVER_PASSWORD"`
}

func NewConfig() *Config {
	path := fetchConfigPath()
	if path == "" {
		panic("config path is empty")
	}

	return NewConfigByPath(path)
}

func NewConfigByPath(configPath string) *Config {

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file not exist: " + configPath)
	}
	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("failed to read config:  " + err.Error())
	}

	// if os.Getenv("DATABASE_HOST") == "db" { // !!! по другому сделать
	// 	cfg.AuthBase = "user:password"
	// }
	// cfg.DatabaseURL = fmt.Sprintf("postgres://%s@%s:5432/%s?%s", cfg.AuthBase, os.Getenv("DATABASE_HOST"), cfg.NameDataBase, cfg.Flags)

	return &cfg
}

// fetchConfigPath fetches config path from command line flag or enviroment variable.
// Priority: flag > env > default
// Default value is empty
func fetchConfigPath() string { //в Посгре немного по другому парситься в целом считываеться с файла
	var configPath string
	flag.StringVar(&configPath, "config", "", "path to config file")
	flag.Parse()
	if configPath == "" {
		configPath = os.Getenv("CONFIG_PATH")
	}
	return configPath
}

// // NewConfig returns a new config instance
// func NewConfig2() *Config {

// 	var configPath string
// 	flag.StringVar(&configPath, "config-path", "configs/local.env", "path to config file")

// 	flag.Parse()
// 	configData, err := os.ReadFile(configPath)
// 	if err != nil {
// 		log.Fatalf("Failed to read config file: %s", err)
// 	}

// 	var config Config

// 	err = yaml.Unmarshal(configData, &config)
// 	if err != nil {
// 		log.Fatalf("Error parsing YAML configuration data %v", err)
// 	}

// 	if err := env.Parse(&config); err != nil {
// 		log.Fatalf("Error parsing environment variables %v", err)
// 	}

// 	if os.Getenv("DATABASE_HOST") == "db" {
// 		config.AuthBase = "user:password"
// 	}
// 	// config.DatabaseURL = fmt.Sprintf("postgres://%s@%s:5432/%s?%s", config.AuthBase, os.Getenv("DATABASE_HOST"), config.NameDataBase, config.Flags)

// 	return &config
// }
