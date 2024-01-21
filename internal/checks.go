package internal

import "bialekredki/card-validator/internal/common"

var majorIndustryIdentifiers = map[int8]string{
	0: "ISO/TC 68",
	1: "Airlines",
	2: "Airlines, financial and other future industry assignments",
	3: "Travel and entertainment",
	4: "Banking and financial",
	5: "Banking and financial",
	6: "Merchandising and banking/financial",
	7: "Petroleum and other future industry assignments",
	8: "Healthcare, telecommunications and other future industry assignments",
	9: "For assignment by national standards bodies",
}

func GetMajorIndustry(cardNumber string) string {
	firstDigit, err := common.FirstDigit(cardNumber)
	if err != nil {
		return "Unknown"
	}
	return majorIndustryIdentifiers[int8(firstDigit)]
}
