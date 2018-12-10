package currencystack

import (
	"currencystack-go/helpers"
	"encoding/json"
	"fmt"
	"strings"
)

// CurrenyConvertionRepository holds the functions for dealing with ip info
type CurrenyConvertionRepository interface {
	GetCurrencyConvertion(base string, target []string) (CurrenyConvertion, error)
}

// CurrenyConvertionAPI implements CurrenyConvertionRepository
type CurrenyConvertionAPI struct {
	HTTP HTTPClient
}

// GetCurrencyConvertion get all the curreny convertions aginst targets ('list of currencies')
func (c CurrenyConvertionAPI) GetCurrencyConvertion(base string, targets []string) (CurrenyConvertion, error) {
	var CurrenyConvertion CurrenyConvertion

	upperBase := strings.ToUpper(base)
	if !helpers.ValidCurrency(upperBase) {
		return CurrenyConvertion, fmt.Errorf("invalid base currency")
	}

	upperTargets := helpers.Map(targets, strings.ToUpper)

	var validCurrencies = helpers.ValidatedCurrencies(upperTargets)
	if len(validCurrencies) == 0 {
		return CurrenyConvertion, fmt.Errorf("invalid target currencies")
	}

	// intialize the currency query params
	opt := Options{Base: upperBase, Target: strings.Join(validCurrencies, ",")}

	data, err := c.HTTP.Get("/currency", opt)

	if err != nil {
		return CurrenyConvertion, err
	}
	err = json.Unmarshal(data, &CurrenyConvertion)
	return CurrenyConvertion, err

}
