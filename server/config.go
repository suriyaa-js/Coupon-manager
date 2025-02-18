package server

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
	"suriyaa.com/coupon-manager/server/props"
)

func LoadConfig(configPath string) (props.Config, error) {
	// Load configuration from file
	fmt.Println("Loading configuration from file: ", configPath)
	data, err := os.ReadFile(configPath)
	if err != nil {
		return props.Config{}, fmt.Errorf("error reading config file: %w", err)
	}

	// Unmarshal the YAML data into the Config struct
	var config props.Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return props.Config{}, fmt.Errorf("error unmarshalling config: %w", err)
	}

	return config, nil
}

