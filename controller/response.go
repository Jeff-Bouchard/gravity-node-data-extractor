package controller

import (
	"encoding/json"
	"fmt"
	"github.com/Gravity-Tech/gravity-node-data-extractor/v2/aggregators"
	extractor "github.com/Gravity-Tech/gravity-node-data-extractor/v2/extractors"
	"github.com/Gravity-Tech/gravity-node-data-extractor/v2/model/ibport"
	//"github.com/Gravity-Tech/gravity-node-data-extractor/v2/model/luport"
	"net/http"

	m "github.com/Gravity-Tech/gravity-node-data-extractor/v2/model"
)

type ResponseController struct {
	TagDelegate *ParamsController
	Config *m.Config
}

func (rc *ResponseController) extractorEnumerator() *extractor.ExtractorEnumerator {
	return extractor.DefaultExtractorEnumerator
}

func (rc *ResponseController) aggregator() aggregators.Aggregator {
	return &aggregators.CommonAggregator{}
}

func (rc *ResponseController) extractor() *extractor.Provider {
	enumerator := rc.extractorEnumerator()

	var impl extractor.IExtractor

	switch enumerator.MatchArgumentEnumeration(rc.TagDelegate.ExtractorType) {
	case enumerator.IBportWavesEth:
		impl = &ibport.IBPortWavesToEthereumExtractor{ Config: rc.Config }
	//case enumerator.LUportWavesEth:
	//	impl = &luport.LUPortWavesToEthereumExtractor{ Config: rc.Config }
	}

	fmt.Printf("Type: %v; Enum: %v \n", rc.TagDelegate.ExtractorType, enumerator.MatchArgumentEnumeration(rc.TagDelegate.ExtractorType))

	return &extractor.Provider{Current: impl}
}

func addBaseHeaders(headers http.Header) {
	headers.Add("Content-Type", "application/json")
}

// swagger:route GET /extracted Extractor getExtractedData
//
// Extracts mapped data
//
// No additional info
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http, https
//
//     Deprecated: false
//
//     Security:
//       api_key:
//
//     Responses:
//       200: BinancePriceIndexResponse
func (rc *ResponseController) GetExtractedData(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		return
	}

	extractor := rc.extractor().Current

	_, extractedData := extractor.Data()

	addBaseHeaders(w.Header())

	b, _ := json.Marshal(extractedData)
	_, _ = fmt.Fprint(w, b)
}

// swagger:route GET /raw Extractor getRawData
//
// Resolves raw data
//
// No additional info
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http, https
//
//     Deprecated: false
//
//     Security:
//       api_key:
//
//     Responses:
//       200: RawData
func (rc *ResponseController) GetRawData(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		return
	}

	extractor := rc.extractor().Current

	rawResponse, _ := extractor.Data()

	addBaseHeaders(w.Header())

	bytes, _ := json.Marshal(&rawResponse)

	_, _ = fmt.Fprint(w, string(bytes))
}

// swagger:route GET /info Extractor getExtractorInfo
//
// Returns extractors common info
//
// No additional info
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http, https
//
//     Deprecated: false
//
//     Security:
//       api_key:
//
//     Responses:
//       200: ExtractorInfo
func (rc *ResponseController) GetExtractorInfo(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		return
	}

	extractor := rc.extractor().Current
	extractorInfo := extractor.Info()

	addBaseHeaders(w.Header())

	bytes, _ := json.Marshal(&extractorInfo)

	_, _ = fmt.Fprint(w, string(bytes))
}

type AggregationRequestBody struct {
	Type   string        `json:"type"`
	Values []interface{} `json:"values"`
}

func (rc *ResponseController) Aggregate(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		return
	}

	type requestType = string
	const (
		typeInt64   requestType = "int64"
		typeFloat64             = "float64"
		typeString              = "string"
	)

	var paramsBody []interface{}

	decoder := json.NewDecoder(req.Body)
	aggregator := rc.aggregator()
	var result interface{}

	addBaseHeaders(w.Header())

	if err := decoder.Decode(&paramsBody); err != nil {
		_, _ = fmt.Fprint(w, fmt.Errorf("Invalid body", err))

		return
	}

	//switch paramsBody.Type {
	//case typeInt64:
	//	result = aggregators.AggregateInt(paramsBody.Values)
	//	break
	//case typeFloat64:
	//	result = aggregators.AggregateFloat(paramsBody.Values)
	//	break
	//case typeString:
	//	result = aggregators.AggregateString(paramsBody.Values)
	//	break
	//}
	result = aggregator.AggregateInt(paramsBody)

	b, _ := json.Marshal(&DataRs{
		Value: result,
	})

	_, _ = fmt.Fprint(w, string(b))
}
