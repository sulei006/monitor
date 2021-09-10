package utils

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"monitor/model"
)

func GetConf(configurePath string) *model.Conf {
	c := &model.Conf{}
	yamlFile, err := ioutil.ReadFile(configurePath)
	if err != nil {
		log.Fatalf("get the db conf fail, error: %s",err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("get the db conf fail, error: %s",err)
	}
	return c
}


