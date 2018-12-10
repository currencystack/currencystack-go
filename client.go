package currencystack

import "fmt"

// Client main client for clear ip
type Client struct {
	HTTPClient   CLHTTPClient
	CurrencyRepo CurrenyConvertionRepository
	BaseURI      string
	APIKey       string
}

// BaseURI base url for clearip
const defaultBaseURI = "https://api.currencystack.io"

// NewClient returns a new ClearIp API client, configured with default HTTPClient.
func NewClient(apiKey string) (*Client, error) {

	if len(apiKey) == 0 {
		return nil, fmt.Errorf("API key required")
	}
	currencyStack := Client{APIKey: apiKey, BaseURI: defaultBaseURI}
	currencyStack.HTTPClient = NewHTTPClient(currencyStack.APIKey, currencyStack.BaseURI)
	currencyStack.setup()
	return &currencyStack, nil
}

func (c *Client) setup() {
	c.CurrencyRepo = CurrenyConvertionAPI{HTTP: c.HTTPClient}

}
