
# Gravity SuSy extractor

[![Build Status](https://drone.gravityhub.org/api/badges/Gravity-Hub-Org/susy-data-extractor/status.svg)](https://drone.gravityhub.org/Gravity-Hub-Org/susy-data-extractor)

## Purpose

The main responsibility of this extractor is
 handling payment state for both source and target chains.
 
## Tests

Internal extractor tests are always run before building public docker image.

To run all **internal** tests manually:

```
go test -v tests/waves-to-eth/*.go
```

## Deployment

