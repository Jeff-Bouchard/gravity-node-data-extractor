package model

// swagger:model
type RawData = byte

type IExtractor interface {
	DataFeedTag() string
	Description() string
	// raw and formated data types
	// first arg should represent type model, second one primitive
	Data() (interface{}, interface{})
	Info() *ExtractorInfo
	extractData(params interface{}) []RawData
	mapData(extractedData []RawData) interface{}
}

// swagger:model
type ExtractorInfo struct {
	Description string `json:"description"`
	DataFeedTag string `json:"datafeedtag"`
}

type ExtractorEnumeration = string
type ExtractorEnumerator struct {
	IBPort_WAVES_ETH, LUPort_WAVES_ETH ExtractorEnumeration
}

var DefaultExtractorEnumerator = &ExtractorEnumerator{
	IBPort_WAVES_ETH: "IB_Port_WAVES_to_ETH",
	LUPort_WAVES_ETH: "LU_Port_WAVES_to_ETH",
}

func (e *ExtractorEnumerator) Default() ExtractorEnumeration {
	return e.IBPort_WAVES_ETH
}

func (e *ExtractorEnumerator) TypeAvailable(enum ExtractorEnumeration) bool {
	available := e.Available()

	for _, item := range available {
		if enum == item { return true }
	}
	return false
}

func (e *ExtractorEnumerator) Available() []ExtractorEnumeration {
	return []string {
		DefaultExtractorEnumerator.IBPort_WAVES_ETH,
		DefaultExtractorEnumerator.LUPort_WAVES_ETH,
	}
}

type ExtractorProvider struct {
	Current IExtractor
}