package utils

import (
	"fmt"
	"strings"
)

// FormatAmountWithCommas formats an integer amount with commas as thousands separators.
// Example: 1234567 becomes "1,234,567"
func FormatAmountWithCommas(amount int64) string {
	amountStr := fmt.Sprintf("%d", amount)
	n := len(amountStr)
	if n <= 3 {
		return amountStr
	}

	var result strings.Builder
	// Handle the first group of digits before the first comma
	firstGroupLen := n % 3
	if firstGroupLen == 0 {
		firstGroupLen = 3 // If divisible by 3, the first group is 3 digits
	}
	result.WriteString(amountStr[:firstGroupLen])
	amountStr = amountStr[firstGroupLen:]

	// Write the rest with commas
	for len(amountStr) > 0 {
		result.WriteString(",")
		result.WriteString(amountStr[:3])
		amountStr = amountStr[3:]
	}

	return result.String()
}

// TranslateCurrency translates common currency codes to Farsi.
func TranslateCurrency(currency string) string {
	switch strings.ToUpper(currency) {
	case "TOMAN":
		return "تومان"
	case "USD":
		return "دلار"
	default:
		return currency // Return original if unknown
	}
}
