package config

import (
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

var (
	config  Config
	configs map[KeyName]string
)

type KeyName = string

type Config struct {
	Server struct {
		Name string `yaml:"name"`
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	}
}

func init() {
	yamlFile, err := ioutil.ReadFile("conf/.default.yaml")
	if err != nil {
		logrus.Fatal(err)
		return
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		logrus.Fatal(err)
		return
	}
	configs = make(map[KeyName]string)
	configs["server.name"] = config.Server.Name
	configs["server.host"] = config.Server.Host
	configs["server.port"] = config.Server.Port
	fmt.Sprint(config)
	fmt.Sprint(configs)
}

func GetString(name KeyName) string {
	return configs[name]
}

func GetInt(name KeyName) int {
	intStr := configs[name]
	if intStr == "" {
		logrus.Warning("Fail to retrieve the configuration with key of: " + name)
	}
	intValue, err := strconv.Atoi(intStr)
	if err != nil {
		logrus.Warning("Fail to convert the configuration value into integer: " + intStr)
	}
	return intValue
}

func GetUint16(name KeyName) uint16 {
	return uint16(GetInt(name))
}
