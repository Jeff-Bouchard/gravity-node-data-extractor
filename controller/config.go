package controller

import (
	"os"
	"fmt"
	"github.com/joho/godotenv"
)

type Config struct {
	SourceSCAddress string
	DestinationSCAddress string
}

func (config *Config) Validate () error {

	values := map[string]string {
		"SourceSCAddress": config.SourceSCAddress,
		"DestinationSCAddress": config.DestinationSCAddress,
	}

	for key, value := range values {
		if value == "" {
			return fmt.Errorf("Error occured. ENV Key is invalid %v \n", key)
		}
	}

	return nil
}

type ConfigBuilder struct {}

func (c *ConfigBuilder) GenerateFromEnvironment () *Config {
	envLoadErr := godotenv.Load(".env")

	if envLoadErr != nil {
		_ = godotenv.Load(".env.example")
	}
	
	return &Config {
		SourceSCAddress: os.Getenv("SOURCE_SC_ADDRESS"),
		DestinationSCAddress: os.Getenv("DESTINATION_SC_ADDRESS"),
	}
}