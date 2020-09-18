package ibport

import (
	waves "github.com/Gravity-Tech/gravity-node-data-extractor/v2/swagger-types/models"
	"strings"
)

/**

# RIDE STATIC KEYS

let FirstRqKey = "first_rq"
let LastRqKey = "last_rq"
let NebulaAddressKey = "nebula_address"

# RIDE DYNAMIC KEYS

func getNextRqKey(id: String) = "next_rq_" + id
func getPrevRqKey(id: String) = "prev_rq_" + id

func getRqReceiverKey(id: String) = "rq_receiver_" + id
func getRqAmountKey(id: String) = "rq_amount_" + id
func getRqStatusKey(id: String) = "rq_status_" + id
func getRqTypeKey(id: String) = "rq_type_" + id

 */

type Hashmap interface {}

/**
#-------------------Constants---------------------------
let NONE = 0

#----Statuses-----
let NEW = 1
let COMPLETED = 2

#----Actions------
let APPROVE = 1
let UNLOCK = 2

#----Types--------
let LOCKTYPE = 1
let UNLOCKTYPE = 2
 */

const (
	firstRqKey = "first_rq"
	lastRqKey = "last_rq"
	nebulaAddressKey = "nebula_address"
)

type RequestID string

type TransferRecord struct {
	RequestID RequestID
	Next, Prev RequestID
	Receiver string
	Amount int
	Status int
	Type int
}

type WavesIBPortHashmap struct {
	hashmap map[RequestID]*TransferRecord
	firstRq, lastRq, nebulaAddress string
}

func (hashmap *WavesIBPortHashmap) ByID (id RequestID) *TransferRecord {
	return hashmap.hashmap[id]
}

func (hashmap *WavesIBPortHashmap) FirstRequest () string {
	return hashmap.firstRq
}

func (hashmap *WavesIBPortHashmap) LastRequest () string {
	return hashmap.lastRq
}

func (hashmap *WavesIBPortHashmap) NebulaAddress () string {
	return hashmap.nebulaAddress
}

func (hashmap *WavesIBPortHashmap) handleDynamicKeyRecord(record *waves.DataEntry) {
	splittedKey := strings.Split(*record.Key, "_")
	requestID := RequestID(splittedKey[len(splittedKey) - 1])
	staticPart := strings.Join(splittedKey[:len(splittedKey) - 2], "_")

	hashmapRecord, ok := hashmap.hashmap[requestID]

	if !ok {
		hashmapRecord = &TransferRecord{}
	}

	switch staticPart {
	case "next_rq_":
		hashmapRecord.Next = record.Value.(string)
	case "prev_rq_":
		hashmapRecord.Prev = record.Value.(string)
	case "rq_receiver_":
		hashmapRecord.Receiver = record.Value.(string)
	case "rq_amount_":
		hashmapRecord.Amount = record.Value.(int)
	case "rq_status_":
		hashmapRecord.Status = record.Value.(int)
	case "rq_type_":
		hashmapRecord.Status = record.Value.(int)
	}

	hashmap.hashmap[requestID] = hashmapRecord
}

func (hashmap *WavesIBPortHashmap) Populate (values []*waves.DataEntry) {

	for _, record := range values {
		switch *record.Key {
		case firstRqKey:
			hashmap.firstRq = record.Value.(string)
		case lastRqKey:
			hashmap.lastRq = record.Value.(string)
		case nebulaAddressKey:
			hashmap.nebulaAddress = record.Value.(string)
		default:
			hashmap.handleDynamicKeyRecord(record)
		}
	}
}
