package configuration

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type promQuery struct {
	expression string `yaml:"expression"`
	start      int    `yaml:"start,omitempty"`
	end        int    `yaml:"end,omitempty"`
}

type configuration struct {
	ocpOauthUrl        string      `yaml:"ocpOauthUrl"`
	prometheusEndpoint string      `yaml:"prometheusEndpoint"`
	username           string      `yaml:"username"`
	password           string      `yaml:"password"`
	queries            []promQuery `yaml:"queries"`
}

func ParseConfigurationFile(filePath string) *configuration {
	configuration := configuration{}
	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error in reading file %s: %v", filePath, err)
	}

	err = yaml.Unmarshal(yamlFile, &configuration)

	if err != nil {
		log.Fatalf("Error in parsing yaml file %s: %v", filePath, err)
	}

	return &configuration
}
