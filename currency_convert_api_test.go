package currencystack

import (
	"io/ioutil"
	"testing"

	"github.com/google/go-querystring/query"
)

func TestGetCurrencyConvertion(t *testing.T) { //                                                                                       url encoded
	http := TestCompanyHTTPClient{t: t, fixtureFilename: "fixtures/CurrenyConvertion.json", expectedURI: "/currency", expectedQuery: "apikey=apiKey&base=EUR&target=USD%2CEGP%2CAED", APIKey: "apiKey", BaseURI: defaultBaseURI}
	api := CurrenyConvertionAPI{HTTP: http}

	CurrenyConvertion, err := api.GetCurrencyConvertion("eur", []string{"Usd", "egp", "aed"})

	if err != nil {
		t.Errorf("Error parsing fixture %s", err)
	}

	if CurrenyConvertion.Base != "EUR" {
		t.Errorf("Name was %s, expected EUR", CurrenyConvertion.Base)
	}
	if CurrenyConvertion.Rates["AED"] != 4.174 {
		t.Errorf("expected AED to equal %f but value was %f", 4.174, CurrenyConvertion.Rates["AED"])
	}

	if CurrenyConvertion.Rates["EGP"] != 20.350 {
		t.Errorf("expected AED to equal %f but value was %f", 20.350, CurrenyConvertion.Rates["AED"])
	}
}

type TestCompanyHTTPClient struct {
	TestHTTPClient
	t               *testing.T
	fixtureFilename string
	expectedURI     string
	expectedQuery   string
	APIKey          string
	BaseURI         string
}

func (t TestCompanyHTTPClient) Get(uri string, queryParams Options) ([]byte, error) {
	queryParams.APIKey = t.APIKey

	if t.expectedURI != uri {
		t.t.Errorf("URI was %s, expected %s", uri, t.expectedURI)
	}

	v, _ := query.Values(queryParams)

	if t.expectedQuery != v.Encode() {
		t.t.Errorf("URI was %s, expected %s", v.Encode(), t.expectedQuery)
	}

	return ioutil.ReadFile(t.fixtureFilename)
}
