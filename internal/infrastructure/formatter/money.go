package formatter

import (
	"bytes"
	"strconv"
)

// FormatMoney formats an integer into a string with dot separators (e.g., 1000 -> 1.000)
func FormatMoney(amount int64) string {
	amountStr := strconv.FormatInt(amount, 10)

	var buf bytes.Buffer

	length := len(amountStr)
	for i, digit := range amountStr {
		buf.WriteRune(digit)

		if (length-i-1)%3 == 0 && i != length-1 {
			buf.WriteRune('.')
		}
	}

	return buf.String()
}
