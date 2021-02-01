package util

// constants for all supported currencies
const (
	USD = "USD"
	PHP = "PHP"
)

// IsSupportedCurrency return true if currency is supported and false otherwise
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, PHP:
		return true
	}
	return false
}
