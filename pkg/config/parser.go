package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

func configParser(filename string) (Options, error) {
	options := Options{}

	configFile, err := os.ReadFile(filename)
	if err != nil {
		return options, fmt.Errorf("%s files Not Found: %v", filename, err)
	}

	err = yaml.Unmarshal(configFile, &options)
	if err != nil {
		return options, fmt.Errorf("%s files is corrupted: %v", filename, err)
	}
	return options, nil
}
