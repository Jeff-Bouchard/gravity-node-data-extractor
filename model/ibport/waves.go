package ibport

import (
	//"bytes"
	"context"
	hexutil "encoding/hex"
	"fmt"
	"github.com/Gravity-Tech/gravity-node-data-extractor/v2/aggregators"
	fetch "github.com/Gravity-Tech/gravity-node-data-extractor/v2/controller/fetch"
	"github.com/Gravity-Tech/gravity-node-data-extractor/v2/extractors"
	"github.com/Gravity-Tech/gravity-node-data-extractor/v2/model/luport"
	eth_client "github.com/Gravity-Tech/gravity-node-data-extractor/v2/client/ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/mr-tron/base58"
	"strconv"
	"strings"
	// "github.com/ethereum/go-ethereum/common/hexutil"
	model "github.com/Gravity-Tech/gravity-node-data-extractor/v2/model"
	//wavesplatform "github.com/wavesplatform/go-lib-crypto"
)

const (
	TransferStatusNew = 1
	TransferStatusCompleted = 2
	// seems like Ethereum only const!
	TransferStatusSuccess = 3
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

func BytesForWAVES(request *luport.TransferRecord) ([]byte, error) {
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

//
//func BytesForETH(request *luport.TransferRecord) ([]byte, error) {
//	var result []byte
//
//	byteBeginning := []byte("c")
//
//	byteRqId := []byte(request.RequestID)
//
//	byteResultStatus := []byte(string(TransferStatusSuccess))
//
//	result = append(result, byteBeginning...)
//	result = append(result, byteRqId...)
//	result = append(result, byteResultStatus...)
//
//	return result, nil
//}

func mapIt(request *eth_client.IBPortRequestCreated) *luport.TransferRecord {
	return &luport.TransferRecord{
		RequestID: luport.RequestID(request.Arg0.String()),
		Next:      "",
		Prev:      "",
		Receiver:  request.Arg1.String(),
		Amount:    0,
		Status:    0,
		Type:      0,
	}
}

func (extractor *IBPortWavesToEthereumExtractor) Data() (interface{}, interface{}) {
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

	currentRequestID := luport.RequestID(wavesRecordHashmap.FirstRequest())

	var newStatusRequests []luport.RequestID

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
			resultString, castError := BytesForWAVES(transferRequest)

			if castError != nil {
				fmt.Printf("Cast error occured: %v \n", castError)
			}

			newStatusRequests = append(newStatusRequests, transferRequest.RequestID)
			resultAn = append(resultAn, resultString...)
		}

		currentRequestID = transferRequest.Next
	}

	ibContract, _ := extractor.ibContract()

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

func (extractor *IBPortWavesToEthereumExtractor) Info() *extractors.ExtractorInfo {
	return &extractors.ExtractorInfo{
		Description: extractor.Description(),
		DataFeedTag: extractor.DataFeedTag(),
	}
}

func (extractor *IBPortWavesToEthereumExtractor) ibContract() (*eth_client.IBPort, error) {
	ethClient := extractor.ethereumClient()
	ibportAddress := extractor.Config.DestinationIBPortAddress

	ibportContract, err := eth_client.NewIBPort(common.HexToAddress(ibportAddress), ethClient)

	if err != nil {
		return nil, err
	}

	return ibportContract, nil
}

func (extractor *IBPortWavesToEthereumExtractor) ethereumClient() *ethclient.Client {
	nodeUrl := extractor.Config.DestinationChainNodeUrl
	var ctx context.Context
	ethClient, err := ethclient.DialContext(ctx, nodeUrl)

	if err != nil {
		return nil
	}

	return ethClient
}

func (extractor *IBPortWavesToEthereumExtractor) wavesClient() *fetch.WavesStateFetcher {
	return &fetch.WavesStateFetcher{ NodeURL: extractor.Config.SourceChainNodeUrl }
}

func (extractor *IBPortWavesToEthereumExtractor) hashmap() *luport.WavesStateHashmap {
	return &luport.WavesStateHashmap{}
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



