package ibport

import (
	"bytes"
	hexutil "encoding/hex"
	"github.com/mr-tron/base58"
	"strings"
	// "github.com/ethereum/go-ethereum/common/hexutil"
	model "github.com/Gravity-Tech/gravity-node-data-extractor/v2/model"
	fetch "github.com/Gravity-Tech/gravity-node-data-extractor/v2/controller/fetch"
	waves "github.com/Gravity-Tech/gravity-node-data-extractor/v2/swagger-types/models"

	wavesplatform "github.com/wavesplatform/go-lib-crypto"

)

type TransferRequestID [32]byte

type IBPortWavesToEthereumExtractor struct {
	Config *model.Config
}

func (extractor *IBPortWavesToEthereumExtractor) DataFeedTag() string {
	return "IBPort_extractor_WAVES_source_ETH_destination"
}

func (extractor *IBPortWavesToEthereumExtractor) Description() string {
	return "This extractor represents IBPort for source chain: WAVES and destination chain: ETH"
}

type transferRequest struct {
	Amount, RequestID, Receiver string
}
func (request *transferRequest) Bytes() ([]byte, error) {
	var result []byte

	byteRqId, err := base58.Decode(request.RequestID)
	if err != nil { return nil, err }
	byteReceiver, err := hexutil.DecodeString(request.Receiver)
	if err != nil { return nil, err }

	result = append(result, byteRqId...)
	result = append(result, byteReceiver...)

	return result, nil
}



func resolveEntry(entries []*waves.DataEntry, key string) *waves.DataEntry {
	for _, entry := range entries {
		if *entry.Key == key {
			return entry
		}
	}

	return nil
}
func filterEntries(entries []*waves.DataEntry, callback func (*waves.DataEntry) bool) []*waves.DataEntry {
	result := make([]*waves.DataEntry, len(entries))

	for _, entry := range entries {
		if callback(entry) {
			result = append(result, entry)
		}
	}

	return result
}

//func IterateEntries(entries []*waves.DataEntry, firstKey string, onNext func(entry *waves.DataEntry)) {
//	// Take first
//	firstEntry := resolveEntry(entries, firstKey)
//	currentRqIdEntry := firstEntry
//
//
//	for {
//		if currentRqIdEntry == nil { break }
//
//		onNext(currentRqIdEntry)
//
//		currentRqIdEntry = resolveEntry(entries, "next_rq_" + currentRqIdEntry.Value.(string))
//	}
//}

func (extractor *IBPortWavesToEthereumExtractor) Data() (interface{}, interface{}) {
	// First iteration
	// Read waves state
	client := extractor.wavesClient()
	addressData, err := client.FetchAddressData(extractor.Config.SourceSCAddress)

	if err != nil {
		return nil, nil
	}

	const (
		TransferStatusNew = 1
		TransferStatusCompleted = 2
	)

	resultAn := make([]byte, len(addressData))
	resultBn := make([]byte, len(addressData))

	//
	// aN computing
	//
	// Taking only new entries - waiting for processing
	newStatusEntries := filterEntries(addressData, func(entry *waves.DataEntry) bool {
		return strings.Contains(*entry.Key, "rq_status_") && entry.Value.(int) == TransferStatusNew
	})

	//
	// Forming transferRequests + mapping to string
	//
	for _, entry := range newStatusEntries {
		entryRqId := strings.Split(*entry.Key, "_")[2]

		amount := resolveEntry(addressData, "rq_amount_" + entryRqId).Value.(string)
		receiver := resolveEntry(addressData, "rq_receiver" + entryRqId).Value.(string)
		resultRequest := &transferRequest{
			Amount:    amount,
			RequestID: entryRqId,
			Receiver:  receiver,
		}

		// EXPLICIT ERROR IGNORE
		resultString, _ := resultRequest.Bytes()

		resultAn = append(resultAn, resultString...)
	}

	//
	// bN computing - ?
	//

	finalResult := append(resultAn, resultBn...)

	return finalResult, finalResult
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

type IBPortWavesToEthereumAggregator struct {
	model.CommonAggregator
}

func (aggregator *IBPortWavesToEthereumAggregator) sort(ls []string) []string {
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
    
    if !sorted { return aggregator.sort(ls) }
    
    return ls
}

func (aggregator *IBPortWavesToEthereumAggregator) hasDuplicates(ls []string) bool {
	hashMap := make(map[int]string, len(ls))

	for index, value := range ls {
		if hashMap[index] != "" {
			return true
		}

		hashMap[index] = value
	}

	return false
}


func (aggregator *IBPortWavesToEthereumAggregator) AggregateString(ls []string) string {
	// var result string

	sorted := aggregator.sort(ls)

	return strings.Join(sorted, "_")
}



