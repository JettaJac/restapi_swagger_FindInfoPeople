package config

import (
	"flag"
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"main/internal/storage/postgre"
	"os"
	"time"
)

type Config struct {
	Env        string           `env:"env" env-default:"local"`
	Namebase   string           `env:"namebase" env-default:"localhost" env-required:"true"`
	HTTPServer HTTPServer       `env-prefix:"http_server."`
	Postgres   postgre.Postgres `env-prefix:"postgres."`
}

type HTTPServer struct {
	Address     string        `env:"address" env-default:":8080"`
	Timeout     time.Duration `env:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `env:"idle_timeout" env-default:"60s"`
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
	fmt.Println(configPath)

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("failed to read config:  " + err.Error())
	}
	fmt.Println(cfg)
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
