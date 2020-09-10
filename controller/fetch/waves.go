package fetch

import (
	"encoding/json"
	"fmt"
	"net/http"
	waves "github.com/Gravity-Tech/gravity-node-data-extractor/v2/swagger-types/models"
)

const nodeUrl = "https://nodes.wavesplatform.com"

type WavesStateFetcher struct {}

func (fetcher *WavesStateFetcher) fetch(path string) (*http.Response, error) {
	return http.Get(nodeUrl + path)
}

func (fetcher *WavesStateFetcher) FetchAddressData(address string) ([]*waves.DataEntry, error) {
	response, err := fetcher.fetch(fmt.Sprintf("/addresses/data/%v", address))

	if err != nil {
		return nil, err
	}

	var result []*waves.DataEntry

	defer response.Body.Close()

	decodeErr := json.NewDecoder(response.Body).Decode(&result)

	if decodeErr != nil { 
		return nil, decodeErr
	}

	return result, nil
}