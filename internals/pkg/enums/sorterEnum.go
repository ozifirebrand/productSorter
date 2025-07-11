package enums

import "strings"

type SorterType string

const (
	SorterTypePriceAsc    SorterType = "PRICE_ASC"
	SorterTypePriceDesc   SorterType = "PRICE_DESC"
	SorterTypeSVRatioAsc  SorterType = "SV_RATIO_ASC"
	SorterTypeSVRatioDesc SorterType = "SV_RATIO_DESC"
)

var validSorterType = map[string]bool{
	string(SorterTypePriceAsc):    true,
	string(SorterTypePriceDesc):   true,
	string(SorterTypeSVRatioAsc):  true,
	string(SorterTypeSVRatioDesc): true,
	// This validation is extendable in the case of adding more sort strategies
}

func IsValidSorterType(input string) bool {
	_, exists := validSorterType[strings.ToUpper(input)]
	return exists
}
