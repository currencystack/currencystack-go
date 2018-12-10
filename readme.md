# clear ip golang library [![Build Status](https://travis-ci.com/currencystack/currencystack-go.svg?branch=master)](https://travis-ci.com/currencystack/currencystack-go)

golang client library for clear ip

## Installation

```bash
go get -u github.com/currencystack/currencystack-go
```

## usage

Get ip info:

```go
package main

import (
	currencystack "github.com/currencystack/currencystack-go"
	"fmt"
)

func main() {

	currencystackClient, err := currencystack.NewClient("API key here")
	if err != nil {
		fmt.Println(err)
		return
	}

	response, err := currencystackClient.CurrencyRepo.GetCurrencyConvertion("eur", []string{"Usd", "egp", "aed"})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(response)
}



```

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)
