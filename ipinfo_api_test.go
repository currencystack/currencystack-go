package currencystack

import (
	"io/ioutil"
	"testing"
)

func TestGetAllDataByIP(t *testing.T) {
	http := TestCompanyHTTPClient{t: t, fixtureFilename: "fixtures/CurrenyConvertion.json", expectedURI: "/ip/192.168.0.1/json", APIKey: "apiKey", BaseURI: defaultBaseURI}
	api := CurrenyConvertionAPI{HTTP: http}

	CurrenyConvertion, err := api.GetAllDataByIP("192.168.0.1")

	if err != nil {
		t.Errorf("Error parsing fixture %s", err)
	}

	if CurrenyConvertion.Country != "Egypt" {
		t.Errorf("Name was %s, expected Egypt", CurrenyConvertion.Country)
	}
}

type TestCompanyHTTPClient struct {
	TestHTTPClient
	t               *testing.T
	fixtureFilename string
	expectedURI     string
	APIKey          string
	BaseURI         string
}

func (t TestCompanyHTTPClient) Get(uri string, queryParams interface{}) ([]byte, error) {
	if t.expectedURI != uri {
		t.t.Errorf("URI was %s, expected %s", uri, t.expectedURI)
	}
	return ioutil.ReadFile(t.fixtureFilename)
}

func (t TestCompanyHTTPClient) Post(uri string, body interface{}) ([]byte, error) {
	if uri != "/companies" {
		t.t.Errorf("Wrong endpoint called")
	}
	return nil, nil
}
