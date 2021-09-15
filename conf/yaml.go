package conf

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type YamlConfig struct {
	Name string `yaml:"name"`
	Host string `yaml:"host"`
	Port int `yaml:"port"`
	App struct {
		Modules []string `yaml:"modules"`
	}
}

func GetConfig(dev string) *YamlConfig  {
	yamlFile, err := ioutil.ReadFile(fmt.Sprintf("./conf/%s.yaml", dev))
	if err != nil {
		panic(err)
	}

	yamlConfig := new(YamlConfig)
	err = yaml.Unmarshal(yamlFile, &yamlConfig)
	if err != nil {
		log.Fatalln("init config failed", err)
	}

	return yamlConfig
}