package server

import (
	"os"

	"gopkg.in/yaml.v2"
)

type config struct {
	DB           string `yaml:"db"`
	ServerPort   string `yaml:"port"`
	ServerSecret string `yaml:"secret"`
}

//returns a new config struct read from yaml file
func NewConfig(path string) config {
	file, err := os.ReadFile(path)
	if err != nil {
		panic("unable to load config data:" + err.Error())
	}
	c := config{}
	err = yaml.Unmarshal(file, &c)
	if err != nil {
		panic("unable to unmarshall config:" + err.Error())
	}

	return c

}
