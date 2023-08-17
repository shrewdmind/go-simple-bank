package util

const (
	USD = "USA"
	EUR = "EUR"
	CAD = "CAD"
)

func IsSuppoertedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, CAD:
		return true
	}
	return false
}