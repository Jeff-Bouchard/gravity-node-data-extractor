package model

import (
	"fmt"
	godotenv "github.com/joho/godotenv"
	"os"
	//"strings"
	"github.com/Gravity-Tech/gravity-node-data-extractor/v2/utils"
	// "strings"
)

type ExtractorType string

const (
	IBPortWavesEth = "ibport-waves-eth"
	LUPortWavesEth = "luport-waves-eth"
	IBPortEthWaves = "ibport-eth-waves"
	LUPortEthWaves = "luport-eth-waves"
)

var availableExtractorPortTypes = []string {
	IBPortWavesEth, LUPortWavesEth, IBPortEthWaves, LUPortEthWaves,
}

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

type ConfigBuilder struct {
	ExtractorType string
}

func (c *ConfigBuilder) GenerateFromEnvironment () *Config {
	extractorType := c.ExtractorType



	if !utils.ContainsString(availableExtractorPortTypes, extractorType) {
		fmt.Errorf("Extractor port type is unavailable: %v \n", extractorType)
		panic(1)
	}

	envLoadErr := godotenv.Load(".env")

	if envLoadErr != nil {
		_ = godotenv.Load(".env.example")
	}
	
	return &Config {
		SourceSCAddress: os.Getenv("SOURCE_SC_ADDRESS"),
		DestinationSCAddress: os.Getenv("DESTINATION_SC_ADDRESS"),
	}
}