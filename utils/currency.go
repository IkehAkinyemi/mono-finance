package utils

// Constants for all supported currencies
const (
	USD = "USD"
	NGN = "NGN"
	GBP = "GBP"
	EUR = "EUR"
)

// IsSupportedCurrency returns true if the currency is supported.
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, NGN, GBP, EUR:
		return true
	}

	return false
}

