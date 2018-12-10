package currencystack

import "encoding/json"

// CurrenyConvertion - represent the convertion data from base currency to targets
type CurrenyConvertion struct {
	Base       string                 `json:"base"`
	LastUpdate string                 `json:"last_update"`
	Status     json.Number            `json:"status"`
	Target     string                 `json:"target"`
	Rates      map[string]interface{} `json:"rates"`
}
