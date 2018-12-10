package helpers

import (
	"github.com/currencystack/currencystack-go/constants"
)

// Intersection - and array of strings with a map
func Intersection(a []string, b map[string]string) (c []string) {
	for _, item := range a {
		if _, ok := b[item]; ok {
			c = append(c, item)
		}
	}
	return
}

// ValidCurrency - return if a one currency key is valid or not
func ValidCurrency(currencyKey string) bool {
	if _, ok := constants.CurrencyKeysList[currencyKey]; ok {
		return true
	}
	return false

}

// ValidatedCurrencies  - return the valid array of the currencies
func ValidatedCurrencies(currencyKeys []string) []string {
	return Intersection(currencyKeys, constants.CurrencyKeysList)
}

// Map - map over an array
func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}
