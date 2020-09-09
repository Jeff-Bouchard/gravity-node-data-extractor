package luport

import (
	// "strings"
	model "github.com/Gravity-Tech/gravity-node-data-extractor/v2/model"
)


type LUPortWavesToEthereumExtractor struct {}

func (extractor *LUPortWavesToEthereumExtractor) DataFeedTag() string {
	return "LUPort_extractor_WAVES_source_ETH_destination"
}

func (extractor *LUPortWavesToEthereumExtractor) Description() string {
	return "This extractor represents LUPort for source chain: WAVES and destination chain: ETH"
}

func (extractor *LUPortWavesToEthereumExtractor) Data() (interface{}, interface{}) {
	return nil, nil
}

func (extractor *LUPortWavesToEthereumExtractor) Info() *model.ExtractorInfo {
	return &model.ExtractorInfo{
		Description: extractor.Description(),
		DataFeedTag: extractor.DataFeedTag(),
	}
}

func (extractor *LUPortWavesToEthereumExtractor) extractData(params interface{}) []model.RawData {
	return make([]model.RawData, 0)
}

func (extractor *LUPortWavesToEthereumExtractor) mapData(extractedData []model.RawData) interface{} {
	return nil
}