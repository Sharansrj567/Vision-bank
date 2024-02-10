package util

// Constants for all supported currencies
const (
	INR = "INR"
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
)

// IsSupportedCurrency returns true if the currency is supported
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case INR, USD, EUR, CAD:
		return true
	}

	return false
}
