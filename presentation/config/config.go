package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Application struct {
		Name string `json:"name"`
		Host string `json:"host"`
		Port string `json:"port"`
	} `json:"application"`
	SupplierConfig struct {
		SjUseMock           bool   `json:"sjUseMock"`
		SjApiBaseUrl        string `json:"sjApiBaseUrl"`
		SjUserName          string `json:"sjUserName"`
		SjUserPassowd       string `json:"sjUserPassowd"`
		SjDeviceId          string `json:"sjDeviceId"`
		SjSubscribeId       string `json:"sjSubscribeId"`
		SjOs                string `json:"sjOs"`
		SjAppsName          string `json:"sjAppsName"`
		SjAppsVersion       string `json:"sjAppsVersion"`
		SjApiSearchPath     string `json:"sjApiSearchPath"`
		SjApiBookPath       string `json:"sjApiBookPath"`
		SjApiBookInfoPath   string `json:"sjApiBookInfoPath"`
		SjApiSetPaymentPath string `json:"sjApiSetPaymentPath"`
	}
}

func LoadConfiguration() (Config, error) {
	var config Config
	configFile, err := os.Open("./config/config.json")
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
		return config, err
	}
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	return config, err
}
