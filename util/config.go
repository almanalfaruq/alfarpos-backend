package util

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Env      string   `yaml:"env"`
	Database Database `yaml:"database"`
}

type Database struct {
	Host       string `yaml:"host"`
	Port       int    `yaml:"port"`
	Username   string `yaml:"username"`
	Password   string `yaml:"password"`
	DBName     string `yaml:"dbname"`
	DBTestName string `yaml:"dbname-test"`
}

type IConfig interface {
	Read(c *Config) error
}

func (config *Config) Read(c *Config) error {
	file, err := ioutil.ReadFile("../config.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(file, &c)
	if err != nil {
		return err
	}
	return nil
}
