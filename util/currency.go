package util

const (
	USD = "USD"
	ERU = "ERU"
	CAD = "CAD"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, ERU, CAD:
		return true
	}
	return false
}
