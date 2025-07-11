package enums

import "strings"

type StrategyType string

const (
	StrategyTypePriceAsc    StrategyType = "PRICE_ASC"
	StrategyTypePriceDesc   StrategyType = "PRICE_DESC"
	StrategyTypeSVRatioAsc  StrategyType = "SV_ASC"
	StrategyTypeSVRatioDesc StrategyType = "SV_DESC"
)

var validStrategyTypes = map[string]bool{
	string(StrategyTypePriceAsc):    true,
	string(StrategyTypePriceDesc):   true,
	string(StrategyTypeSVRatioAsc):  true,
	string(StrategyTypeSVRatioDesc): true,
	// This validation is extendable in the case of adding more sort strategies
}

func IsValidStrategyType(input string) bool {
	_, exists := validStrategyTypes[strings.ToUpper(input)]
	return exists
}
