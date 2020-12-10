package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Config struct {
	db DBConfig
}
type DBConfig struct {
	name     string
	addr     string
	username string
	password string
}

// 单例
var CFG Config
var DB_CONFIG DBConfig

func (cfg *Config) Init() {
	CFG = GetConfig()
	DB_CONFIG = GetConfig().db
}

func GetConfig() Config {
	file, err := os.Open("config.yaml")
	if err != nil {
		panic(err)
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	cfg := Config{}
	err = yaml.Unmarshal(bytes, &cfg)
	if err != nil {
		panic(err)
	}
	fmt.Println(cfg)
	return cfg
}
