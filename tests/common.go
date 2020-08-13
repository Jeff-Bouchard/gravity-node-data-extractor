package tests


type InternalAggregationRequest struct {
	RequestID string
	Receiver string
	Amount int64
}

type InternalAggregationRequestList struct {
	Values []*InternalAggregationRequest
}


type SwapRequest struct {
	Sender, Recipient string
	Amount int64
	Currency string
}

type ExtractionReadyState struct {
	ProcessedRequestIDList []string
	BlockForRequest int64
	BlockInterestRange []int64
	ReceivingAddress, SenderAddress string

	SwapCurrency string
}

func (state *ExtractionReadyState) HasBlockInterest (height int64) bool {
	begin, end := state.BlockInterestRange[0], state.BlockInterestRange[1]

	return height >= begin && height <= end
}
