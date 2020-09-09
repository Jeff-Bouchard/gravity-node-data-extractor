package ibport

import (
	model "github.com/Gravity-Tech/gravity-node-data-extractor/v2/model"
)


// func fetchState

type IBPortWavesToEthereumExtractor struct {}

func (extractor *IBPortWavesToEthereumExtractor) DataFeedTag() string {
	return "IBPort_extractor_WAVES_source_ETH_destination"
}

func (extractor *IBPortWavesToEthereumExtractor) Description() string {
	return "This extractor represents IBPort for source chain: WAVES and destination chain: ETH"
}

func (extractor *IBPortWavesToEthereumExtractor) Data() (interface{}, interface{}) {
	return nil, nil
}

func (extractor *IBPortWavesToEthereumExtractor) Info() *model.ExtractorInfo {
	return &model.ExtractorInfo{
		Description: extractor.Description(),
		DataFeedTag: extractor.DataFeedTag(),
	}
}

func (extractor *IBPortWavesToEthereumExtractor) extractData() []model.RawData {
	return make([]model.RawData, 0)
}

func (extractor *IBPortWavesToEthereumExtractor) mapData() interface{} {
	return nil
}

type IBPortWavesToEthereumAggregator struct {
	model.CommonAggregator
}

func (fetcher *IBPortWavesToEthereumAggregator) AggregateString([]string list) string {
	var result string

	return result
}



