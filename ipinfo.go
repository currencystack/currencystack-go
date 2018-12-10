package currencystack

import "encoding/json"

// CurrenyConvertion - represent the full ip data from the get response
type CurrenyConvertion struct {
	Base       string      `json:"base"`
	LastUpdate string      `json:"last_update"`
	Status     json.Number `json:"status"`
	Target     string      `json:"target"`
	Rates      interface{} `json:"rates"`
}
