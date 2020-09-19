package extractors

import "github.com/Gravity-Tech/gravity-node-data-extractor/v2/model"

var DefaultExtractorEnumerator = &ExtractorEnumerator{
	IBportWavesEth: "IB_Port_WAVES_ETH",
	LUportWavesEth: "LU_Port_WAVES_ETH",
	IBportEthWaves: "IB_Port_ETH_WAVES",
	LUportEthWaves: "LU_Port_ETH_WAVES",
}

func (e *ExtractorEnumerator) Default() ExtractorEnumeration {
	return e.LUportWavesEth
}

func (e *ExtractorEnumerator) TypeAvailable(enum ExtractorEnumeration) bool {
	available := e.Available()

	for _, item := range available {
		if enum == item { return true }
	}
	return false
}

func (e *ExtractorEnumerator) Available() []ExtractorEnumeration {
	return []string {
		DefaultExtractorEnumerator.IBportWavesEth,
		DefaultExtractorEnumerator.LUportWavesEth,
		DefaultExtractorEnumerator.IBportEthWaves,
		DefaultExtractorEnumerator.LUportEthWaves,
	}
}

func (e *ExtractorEnumerator) MatchArgumentEnumeration (enumeration ExtractorEnumeration) ExtractorEnumeration {
	switch enumeration {
	case model.LUPortEthWaves:
		return DefaultExtractorEnumerator.LUportEthWaves
	case model.LUPortWavesEth:
		return DefaultExtractorEnumerator.LUportWavesEth
	case model.IBPortEthWaves:
		return DefaultExtractorEnumerator.IBportEthWaves
	case model.IBPortWavesEth:
		return DefaultExtractorEnumerator.IBportWavesEth
	}

	return ""
}

type Provider struct {
	Current IExtractor
}