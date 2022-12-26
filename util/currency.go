package util

const (
	USD = "USD"
	EUR = "EUR"
	EGP = "EGP"
	CAD = "CAD"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EGP, EUR, CAD:
		return true
	}
	return false
}
