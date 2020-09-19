package luport

import (
	//"bytes"
	hexutil "encoding/hex"
	"fmt"
	"github.com/Gravity-Tech/gravity-node-data-extractor/v2/aggregators"
	fetch "github.com/Gravity-Tech/gravity-node-data-extractor/v2/controller/fetch"
	"github.com/Gravity-Tech/gravity-node-data-extractor/v2/extractors"
	"github.com/mr-tron/base58"
	"strconv"
	"strings"
	// "github.com/ethereum/go-ethereum/common/hexutil"
	model "github.com/Gravity-Tech/gravity-node-data-extractor/v2/model"
	waves "github.com/Gravity-Tech/gravity-node-data-extractor/v2/swagger-types/models"

	//wavesplatform "github.com/wavesplatform/go-lib-crypto"

)

const (
	TransferStatusNew = 1
	TransferStatusCompleted = 2

	// seems like Ethereum only const!
	TransferStatusSuccess = 3
)

type TransferRequestID [32]byte

type LUPortWavesToEthereumExtractor struct {
	Config *model.Config
}

func (extractor *LUPortWavesToEthereumExtractor) DataFeedTag() string {
	return "IBPort_extractor_WAVES_source_ETH_destination"
}

func (extractor *LUPortWavesToEthereumExtractor) Description() string {
	return "This extractor represents IBPort for source chain: WAVES and destination chain: ETH"
}

func (request *TransferRecord) BytesForWAVES() ([]byte, error) {
	result := make([]byte, 0)

	byteBeginning := []byte("m")

	byteRqId, err := base58.Decode(string(request.RequestID))
	if err != nil {
		return nil, err
	}

	byteAmount := []byte(strconv.FormatInt(request.Amount, 10))

	requestReceiver := request.Receiver

	if strings.HasPrefix(requestReceiver, "0x") {
		requestReceiver = requestReceiver[2:]
	}

	byteReceiver, err := hexutil.DecodeString(requestReceiver)
	if err != nil { return nil, err }

	result = append(result, byteBeginning...)
	result = append(result, byteRqId...)
	result = append(result, byteAmount...)
	result = append(result, byteReceiver...)

	return result, nil
}

func (request *TransferRecord) BytesForETH() ([]byte, error) {
	var result []byte

	byteBeginning := []byte("c")

	byteRqId := []byte(request.RequestID)

	byteResultStatus := []byte(string(TransferStatusSuccess))

	result = append(result, byteBeginning...)
	result = append(result, byteRqId...)
	result = append(result, byteResultStatus...)

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

func (extractor *LUPortWavesToEthereumExtractor) Data() (interface{}, interface{}) {
	// First iteration
	// Read waves state
	client := extractor.wavesClient()

	addressData, err := client.FetchAddressData(extractor.Config.SourceLUPortAddress)

	if err != nil {
		fmt.Printf("Error occured: %v \n", err)
		return nil, nil
	}


	resultAn := make([]byte, len(addressData))
	//resultBn := make([]byte, len(addressData))


	// Populate
	//
	// aN: Forming transferRequests + mapping to string
	//
	wavesRecordHashmap := extractor.hashmap()
	wavesRecordHashmap.Populate(addressData)

	currentRequestID := RequestID(wavesRecordHashmap.FirstRequest())

	var newStatusRequests []RequestID

	for {
		if currentRequestID == "" { break }

		transferRequest := wavesRecordHashmap.ByID(currentRequestID)

		fmt.Printf("TransferRequest: %+v \n", transferRequest)

		//
		// aN: Computing
		//
		// Taking only new entries - waiting for processing
		//
		if transferRequest.Status == TransferStatusNew {
			//entryRqId := strings.Split(*entry.Key, "_")[2]

			fmt.Printf("GOT NEW REQUEST: %v \n", transferRequest)
			// EXPLICIT ERROR IGNORE
			resultString, castError := transferRequest.BytesForWAVES()

			if castError != nil {
				fmt.Printf("Cast error occured: %v \n", castError)
			}

			newStatusRequests = append(newStatusRequests, transferRequest.RequestID)
			resultAn = append(resultAn, resultString...)
		}

		currentRequestID = transferRequest.Next
	}


	//
	////
	//// bN: Computing - ?
	////
	//
	////
	//// bN: Forming transferRequests + mapping to string
	////
	//// WAVES ENTRIES ARE JUST FOR EXAMPLE
	//for _, entryID := range newStatusRequests {
	//	//entryRqId := strings.Split(*entry.Key, "_")[2]
	//	//
	//	//amount := resolveEntry(addressData, "rq_amount_" + entryRqId).Value.(string)
	//	//receiver := resolveEntry(addressData, "rq_receiver" + entryRqId).Value.(string)
	//	//status := resolveEntry(addressData, "rq_status" + entryRqId).Value.(string)
	//	instance := wavesRecordHashmap.ByID(entryID)
	//
	//	resultRequest := &transferRequest{
	//		Amount:    amount,
	//		RequestID: entryRqId,
	//		Receiver:  receiver,
	//	}
	//
	//	if status == string(TransferStatusCompleted) {
	//		// EXPLICIT ERROR IGNORE
	//		resultString, _ := resultRequest.BytesForETH()
	//
	//		resultBn = append(resultBn, resultString...)
	//	}
	//}
	//
	//finalResult := append(resultAn, resultBn...)
	//
	//return finalResult, finalResult

	fmt.Printf("Result AN: %v \n", resultAn)

	return resultAn, resultAn
}

func (extractor *LUPortWavesToEthereumExtractor) Info() *extractors.ExtractorInfo {
	return &extractors.ExtractorInfo{
		Description: extractor.Description(),
		DataFeedTag: extractor.DataFeedTag(),
	}
}

func (extractor *LUPortWavesToEthereumExtractor) ethereumClient() interface{} {
	return nil
}

func (extractor *LUPortWavesToEthereumExtractor) wavesClient() *fetch.WavesStateFetcher {
	return &fetch.WavesStateFetcher{ NodeURL: extractor.Config.SourceChainNodeUrl }
}

func (extractor *LUPortWavesToEthereumExtractor) hashmap() *WavesLUPortHashmap {
	return &WavesLUPortHashmap{}
}

type IBPortWavesToEthereumAggregator struct {
	aggregators.CommonAggregator
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



