package config

import (
	"os"
	"sync"

	"github.com/pelletier/go-toml/v2"
)

var (
	config Config
	once   sync.Once
)

type Config struct {
	Title      string `toml:"title"`
	DebugModel bool   `toml:"debug"`
	Prometheus struct {
		Port int `toml:"port"`
	} `toml:"prometheus"`
	JsServer struct {
		Host string `toml:"host"`
	} `toml:"js_server"`
	Server struct {
		Port        int `toml:"port"`
		MaxPageSize int `toml:"max_page_size"`
		Cors        struct {
			AllowedOrigins []string `toml:"allowed_origins"`
			AllowedHeaders []string `toml:"allowed_headers"`
		} `toml:"cors"`
	} `toml:"server"`
	DataBase struct {
		Type string `toml:"type"`
		DSN  struct {
			Addr           string `toml:"addr"`
			DB             string `toml:"db"`
			Username       string `toml:"username"`
			Password       string `toml:"password"`
			ConnectTimeout int    `toml:"connect_timeout"`
			MaxIdleConns   int    `toml:"max_idle_conns"`
		} `toml:"dsn"`
	} `toml:"database"`
	MinIO struct {
		ADDR            string `toml:"addr"`
		AccessKeyID     string `toml:"access_key_id"`
		SecretAccessKey string `toml:"secret_access_key"`
		SSL             bool   `toml:"ssl"`
	} `toml:"minio"`
}

func GetConfig() *Config {
	once.Do(func() {
		tomlData, err := os.ReadFile("config.toml")
		if err != nil {
			panic(err)
		}
		if err := toml.Unmarshal(tomlData, &config); err != nil {
			panic(err)
		}
	})
	return &config
}
