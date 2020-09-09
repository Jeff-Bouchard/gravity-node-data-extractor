package ibport

import (
	"strings"
	model "github.com/Gravity-Tech/gravity-node-data-extractor/v2/model"
)


// func fetchState


// type IExtractor interface {
// 	DataFeedTag() string
// 	Description() string
// 	Data() (interface{}, interface{})
// 	Info() *ExtractorInfo
// 	extractData(params interface{}) []RawData
// 	mapData(extractedData []RawData) interface{}
// }

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

func (extractor *IBPortWavesToEthereumExtractor) extractData(params interface{}) []model.RawData {
	return make([]model.RawData, 0)
}

func (extractor *IBPortWavesToEthereumExtractor) mapData(extractedData []model.RawData) interface{} {
	return nil
}

type IBPortWavesToEthereumAggregator struct {
	model.CommonAggregator
}

func (fetcher *IBPortWavesToEthereumAggregator) sort(ls []string) []string {
	sorted := true
    for index, value := range ls {
        
        if index == len(ls) - 1 { break }
        
        next := ls[index + 1]
        if value > next {
            ls[index] = next
            ls[index+1] = value
            sorted = false
        }
    }
    
    if !sorted { return fetcher.sort(ls) }
    
    return ls
}

func (fetcher *IBPortWavesToEthereumAggregator) hasDuplicates(ls []string) bool {
	hashMap := make(map[int]string, len(ls))

	for index, value := range ls {
		if hashMap[index] != "" {
			return true
		}

		hashMap[index] = value
	}

	return false
}


func (fetcher *IBPortWavesToEthereumAggregator) AggregateString(ls []string) string {
	// var result string

	sorted := fetcher.sort(ls)

	return strings.Join(sorted, "_")
}



