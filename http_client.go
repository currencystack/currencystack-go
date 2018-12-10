package currencystack

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
)

// HTTPClient -
type HTTPClient interface {
	Get(string, Options) ([]byte, error)
}

// CLHTTPClient - impelemts HTTPClient interface clear ip default client
type CLHTTPClient struct {
	*http.Client
	BaseURI string
	APIKey  string
}

//NewHTTPClient - create the clear ip client with api key
func NewHTTPClient(apiKey string, BaseURI string) CLHTTPClient {
	return CLHTTPClient{Client: &http.Client{
		Timeout: time.Second * 10,
	},
		APIKey:  apiKey,
		BaseURI: BaseURI,
	}
}

// Options - queryParams schema
type Options struct {
	APIKey string `url:"apikey"`
	Base   string `url:"base"`
	Target string `url:"target"`
}

// Get default http client
func (c CLHTTPClient) Get(url string, queryParams Options) ([]byte, error) {
	// Setup request
	req, _ := http.NewRequest("GET", c.BaseURI+url, nil)
	req.Header.Add("Accept", "application/json")
	queryParams.APIKey = c.APIKey

	addQueryParams(req, queryParams)

	// Do request
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read response
	data, err := c.readAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 400 {
		return nil, c.parseResponseError(data, resp.StatusCode)
	}
	return data, err
}

func (c CLHTTPClient) parseResponseError(data []byte, statusCode int) HTTPError {
	errorList := HTTPErrorList{}
	err := json.Unmarshal(data, &errorList)
	if err != nil {
		return NewUnknownHTTPError(statusCode)
	}
	if len(errorList.Errors) == 0 {
		return NewUnknownHTTPError(statusCode)
	}
	httpError := errorList.Errors[0]
	httpError.StatusCode = statusCode
	return httpError // only care about the first
}

func (c CLHTTPClient) readAll(body io.Reader) ([]byte, error) {
	b, err := ioutil.ReadAll(body)

	return b, err
}

func addQueryParams(req *http.Request, params Options) {
	v, _ := query.Values(params)
	req.URL.RawQuery = v.Encode()
}
