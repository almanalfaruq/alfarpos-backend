package util

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Env         string      `yaml:"env"`
	SecretKey   string      `yaml:"secret-key"`
	Database    Database    `yaml:"database"`
	ShopProfile ShopProfile `yaml:"shop-profile"`
	Log         Log         `yaml:"log"`
}

type Database struct {
	Host       string `yaml:"host"`
	Port       int    `yaml:"port"`
	Username   string `yaml:"username"`
	Password   string `yaml:"password"`
	DBName     string `yaml:"dbname"`
	DBTestName string `yaml:"dbname-test"`
}

type ShopProfile struct {
	Name            string `yaml:"name"`
	Description     string `yaml:"description"`
	Address         string `yaml:"address"`
	Phone           string `yaml:"phone"`
	NPWP            string `yaml:"npwp"`
	ThankyouMessage string `yaml:"thankyou_message"`
	FootNote        string `yaml:"foot_note"`
}

type Log struct {
	PathInfo  string `yaml:"path_info"`
	PathDebug string `yaml:"path_debug"`
}

type IConfig interface {
	Read(filePath string, c *Config) error
}

func (config *Config) Read(filePath string, c *Config) error {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(file, &c)
	if err != nil {
		return err
	}
	return nil
}
