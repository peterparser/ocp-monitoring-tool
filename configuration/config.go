package configuration

import (
	"io/ioutil"
	"log"

	"../util"
	"gopkg.in/yaml.v3"
)

func ParseConfigurationFile(filePath string) *util.Configuration {
	configuration := util.Configuration{}
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
