package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

const configFile = "config.yml"

func LoadConfig() error {
	source, err := ioutil.ReadFile(configFile)
	if err != nil {
		fmt.Errorf(" load file error %v ", err)
		return err
	}
	if err := yaml.Unmarshal(source, &Config); err != nil {
		fmt.Errorf(" Error %v ", err)
		return err
	}
	return nil
}
func InitApi() []CryptoApi {
	var apis []CryptoApi
	for _, c := range Config.Api {
		apis = append(apis, &CryptoCompare{Config: c})
	}
	return apis
}
