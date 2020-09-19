package model

import (
	"fmt"
	godotenv "github.com/joho/godotenv"
	"os"
	//"strings"
	//"github.com/Gravity-Tech/gravity-node-data-extractor/v2/utils"
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
	SourceChainNodeUrl string
	DestinationChainNodeUrl string

	SourceLUPortAddress string
	SourceIBPortAddress string
	DestinationLUPortAddress string
	DestinationIBPortAddress string
}

func (config *Config) Validate () error {

	values := map[string]string {
		"SourceLUPortAddress": config.SourceLUPortAddress,
		"SourceIBPortAddress": config.SourceIBPortAddress,
		"DestinationLUPortAddress": config.DestinationLUPortAddress,
		"DestinationIBPortAddress": config.DestinationIBPortAddress,
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

	isAvailable := false
	for _, availableType := range availableExtractorPortTypes {
		if availableType == extractorType {
			isAvailable = true
		}
	}
	if !isAvailable {
		fmt.Errorf("Extractor port type is unavailable: %v \n", extractorType)
		panic(1)
	}

	envLoadErr := godotenv.Load(".env")

	if envLoadErr != nil {
		_ = godotenv.Load(".env.example")
	}
	
	return &Config {
		SourceChainNodeUrl: os.Getenv("SOURCE_CHAIN_PUBLIC_NODE"),
		DestinationChainNodeUrl: os.Getenv("DESTINATION_CHAIN_PUBLIC_NODE"),
		SourceLUPortAddress: os.Getenv("SOURCE_LU_PORT"),
		SourceIBPortAddress: os.Getenv("SOURCE_IB_PORT"),
		DestinationLUPortAddress: os.Getenv("DESTINATION_LU_PORT"),
		DestinationIBPortAddress: os.Getenv("DESTINATION_IB_PORT"),
	}
}