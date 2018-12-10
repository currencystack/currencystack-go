package currencystack

type TestHTTPClient struct{}

func (h TestHTTPClient) Get(uri string, queryParams interface{}) ([]byte, error) { return nil, nil }
