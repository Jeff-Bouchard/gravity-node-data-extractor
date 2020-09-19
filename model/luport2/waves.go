package luport2

import (
	// "strings"
	"github.com/Gravity-Tech/gravity-node-data-extractor/v2/extractors"
	model "github.com/Gravity-Tech/gravity-node-data-extractor/v2/model"
)


type LUPortWavesToEthereumExtractor struct {
	Config *model.Config
}

func (extractor *LUPortWavesToEthereumExtractor) DataFeedTag() string {
	return "LUPort_extractor_WAVES_source_ETH_destination"
}

func (extractor *LUPortWavesToEthereumExtractor) Description() string {
	return "This extractor represents LUPort for source chain: WAVES and destination chain: ETH"
}

func (extractor *LUPortWavesToEthereumExtractor) Data() (interface{}, interface{}) {
	return nil, nil
}

func (extractor *LUPortWavesToEthereumExtractor) Info() *extractors.ExtractorInfo {
	return &extractors.ExtractorInfo{
		Description: extractor.Description(),
		DataFeedTag: extractor.DataFeedTag(),
	}
}

func (extractor *LUPortWavesToEthereumExtractor) extractData(params interface{}) []extractors.RawData {
	return make([]extractors.RawData, 0)
}

func (extractor *LUPortWavesToEthereumExtractor) mapData(extractedData []extractors.RawData) interface{} {
	return nil
}