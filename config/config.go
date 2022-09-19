package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

type Config struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

var config Config

func init() {
	env := os.Getenv("NX_ENV")
	if env == "" {
		env = "development"
	}

	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	configFile, err := ioutil.ReadFile(filepath.Join(home, ".nx", "config.yml"))

	data := make(map[string]Config)

	err2 := yaml.Unmarshal(configFile, data)
	if err2 != nil {
		panic(err2)
	}

	if val, ok := data[env]; ok {
		config = val
	} else {
		panic("No configured environment for " + env)
	}
}

func Host() string {
	if config.Host != "" {
		return config.Host
	} else {
		return "localhost"
	}
}

func Port() string {
	return strconv.Itoa(config.Port)
}
