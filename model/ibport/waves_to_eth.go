package ibport

import (
	"strings"
	// "github.com/ethereum/go-ethereum/common/hexutil"
	model "github.com/Gravity-Tech/gravity-node-data-extractor/v2/model"
	fetch "github.com/Gravity-Tech/gravity-node-data-extractor/v2/controller/fetch"
	waves "github.com/Gravity-Tech/gravity-node-data-extractor/v2/swagger-types/models"
)

type IBPortWavesToEthereumExtractor struct {
	Config *model.Config
}

func (extractor *IBPortWavesToEthereumExtractor) DataFeedTag() string {
	return "IBPort_extractor_WAVES_source_ETH_destination"
}

func (extractor *IBPortWavesToEthereumExtractor) Description() string {
	return "This extractor represents IBPort for source chain: WAVES and destination chain: ETH"
}


func resolveEntry(entries []*waves.DataEntry, key string) *waves.DataEntry {
	for _, entry := range entries {
		if *entry.Key == key {
			return entry
		}
	}

	return nil
}
func (extractor *IBPortWavesToEthereumExtractor) Data() (interface{}, interface{}) {

	// First iteration
	// Read waves state
	client := extractor.wavesClient()
	addressDataCollection, err := client.FetchAddressData(extractor.Config.SourceSCAddress)

	if err != nil {
		return nil, nil
	}

	// type wavesTransferRequest struct {
	// 	Amount uint64
	// 	Receiver, RequestId string
	// }

	transferRequestID := resolveEntry(addressDataCollection, "first")
	transferAmount := resolveEntry(addressDataCollection, "rq_amount_" + transferRequestID.Value.(string))
	requestReceiver := resolveEntry(addressDataCollection, "rq_receiver_" + transferRequestID.Value.(string))

	// m{base58ToBytes(rqId)}{toBytes32(amount)}{HextoBytes20(reciver)}

	resultString := strings.Join(
		[]string {
			"m",
			base58ToBytes(transferRequestID.Value.(string)),
			toBytes32(transferAmount.Value.(string)),
			hexToBytes(requestReceiver.Value.(string)),
		},
		"",
	)

	return resultString, resultString
}

func base58ToBytes(rqId string) string {
	return ""
}
func toBytes32(amount string) string {
	return ""
}
func hexToBytes(receiver string) string {
	return ""
}


func (extractor *IBPortWavesToEthereumExtractor) Info() *model.ExtractorInfo {
	return &model.ExtractorInfo{
		Description: extractor.Description(),
		DataFeedTag: extractor.DataFeedTag(),
	}
}

func (extractor *IBPortWavesToEthereumExtractor) ethereumClient() interface{} {
	return nil
}

func (extractor *IBPortWavesToEthereumExtractor) wavesClient() *fetch.WavesStateFetcher {
	return &fetch.WavesStateFetcher{}
}

// func (extractor *IBPortWavesToEthereumExtractor) extractData(params interface{}) []model.RawData {
// 	return make([]model.RawData, 0)
// }

// func (extractor *IBPortWavesToEthereumExtractor) mapData(extractedData []model.RawData) interface{} {
// 	return nil
// }

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



