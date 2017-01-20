package config

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	LogServer struct {
		LogAddress string `yaml:"log_address"`
		LogPort    int    `yaml:"log_port"`
	} `yaml:"log_server"`

	Database struct {
		DSN string `yaml:"dsn"`
	} `yaml:"database"`
}

var Conf Config

func InitialiseConfiguration() {
	configuration, _ := ioutil.ReadFile("./config.yml")
	err := yaml.Unmarshal(configuration, &Conf)

	if err != nil {
		log.Println("Failed to initialise configuration:", err)
	}
}
