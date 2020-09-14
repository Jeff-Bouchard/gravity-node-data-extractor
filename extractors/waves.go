package extractors

var DefaultExtractorEnumerator = &ExtractorEnumerator{
	IbportWavesEth: "IB_Port_WAVES_to_ETH",
	LuportWavesEth: "LU_Port_WAVES_to_ETH",
}

func (e *ExtractorEnumerator) Default() ExtractorEnumeration {
	return e.IbportWavesEth
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
		DefaultExtractorEnumerator.IbportWavesEth,
		DefaultExtractorEnumerator.LuportWavesEth,
	}
}

func (e *ExtractorEnumerator) MatchArgumentEnumeration (enumeration ExtractorEnumeration) ExtractorEnumeration {
	switch enumeration {
	case "ib-waves-eth":
		return DefaultExtractorEnumerator.IbportWavesEth
	case "lu-waves-eth":
		return DefaultExtractorEnumerator.LuportWavesEth
	}

	return ""
}

type Provider struct {
	Current IExtractor
}