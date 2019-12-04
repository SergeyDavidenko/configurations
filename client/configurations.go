package client

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// Config ...
var Config *ConfigData

// GetConfigFrom ...
func GetConfigFrom(url string, login string, password string) (*ConfigData, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(login, password)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &Config); err != nil {
		return nil, err
	}
	return Config, nil
}

// GetBoolFromConfig ...
func GetBoolFromConfig(config *ConfigData, property string) (bool, error) {
	for _, element := range config.PropertySources {
		_, ok := element.Source[property]
		if ok {
			prop, err := strconv.ParseBool(element.Source[strings.ToLower(property)])
			if err != nil {
				return false, err
			}
			return prop, nil
		}
	}
	return false, errors.New("Error to get element")
}

// GetStringFromConfig ...
func GetStringFromConfig(config *ConfigData, property string) (string, error) {
	for _, element := range config.PropertySources {
		_, ok := element.Source[property]
		if ok {
			return element.Source[property], nil
		}
	}
	return "", errors.New("Error to get element")
}

// GetIntFromConfig ...
func GetIntFromConfig(config *ConfigData, property string) (int, error) {
	for _, element := range config.PropertySources {
		_, ok := element.Source[property]
		if ok {
			prop, err := strconv.Atoi(element.Source[property])
			if err != nil {
				return 0, err
			}
			return prop, nil
		}
	}
	return 0, errors.New("Error to get element")
}

// GetFloat64FromConfig ...
func GetFloat64FromConfig(config *ConfigData, property string) (float64, error) {
	for _, element := range config.PropertySources {
		_, ok := element.Source[property]
		if ok {
			prop, err := strconv.ParseFloat(element.Source[property], 64)
			if err != nil {
				return 0, err
			}
			return prop, nil
		}
	}
	return 0, errors.New("Error to get element")
}

// GetFloat32FromConfig ...
func GetFloat32FromConfig(config *ConfigData, property string) (float32, error) {
	for _, element := range config.PropertySources {
		_, ok := element.Source[property]
		if ok {
			prop, err := strconv.ParseFloat(element.Source[property], 32)
			if err != nil {
				return 0, err
			}
			return float32(prop), nil
		}
	}
	return 0, errors.New("Error to get element")
}

// GetStringOrDefault ..
func GetStringOrDefault(property string, def string) string {
	prop, err := GetStringFromConfig(Config, property)
	if err != nil {
		return def
	}
	return prop
}

// GetIntOrDefault ..
func GetIntOrDefault(property string, def int) int {
	prop, err := GetIntFromConfig(Config, property)
	if err != nil {
		return def
	}
	return prop
}

// GetFloat32OrDefault ..
func GetFloat32OrDefault(property string, def float32) float32 {
	prop, err := GetFloat32FromConfig(Config, property)
	if err != nil {
		return def
	}
	return prop
}

// GetFloat64OrDefault ..
func GetFloat64OrDefault(property string, def float64) float64 {
	prop, err := GetFloat64FromConfig(Config, property)
	if err != nil {
		return def
	}
	return prop
}

// GetBoolOrDefault ..
func GetBoolOrDefault(property string, def bool) bool {
	prop, err := GetBoolFromConfig(Config, property)
	if err != nil {
		return def
	}
	return prop
}
