package fetch

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	waves "github.com/Gravity-Tech/gravity-node-data-extractor/v2/swagger-types/models"
)

type WavesStateFetcher struct {
	NodeURL string
}

func (fetcher *WavesStateFetcher) fetch(path string) (*http.Response, error) {
	return http.Get(fetcher.NodeURL + path)
}

func (fetcher *WavesStateFetcher) FetchAddressData(address string) ([]*waves.DataEntry, error) {
	requestURL := fmt.Sprintf("/addresses/data/%v", address)
	response, err := fetcher.fetch(requestURL)

	if err != nil {
		return nil, err
	}

	var result []*waves.DataEntry

	defer response.Body.Close()

	byteValue, _ := ioutil.ReadAll(response.Body)

	decodeErr := json.Unmarshal(byteValue, &result)

	if decodeErr != nil { 
		return nil, decodeErr
	}

	return result, nil
}