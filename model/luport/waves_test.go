package luport

import (
	"github.com/Gravity-Tech/gravity-node-data-extractor/v2/model"
	"os"
	"testing"
)

var extractorConfig *model.Config

func prepare() {
	configBuiler := &model.ConfigBuilder{ ExtractorType: model.LUPortWavesEth }
	extractorConfig = configBuiler.GenerateFromEnvironment()
}

func TestConfig(t *testing.T) {
	validationErr := extractorConfig.Validate()

	if validationErr != nil {
		t.Error(validationErr)
		t.Fail()
	}
}

func TestMain(m *testing.M) {
	prepare()

	// call flag.Parse() here if TestMain uses flags
	os.Exit(m.Run())
}
