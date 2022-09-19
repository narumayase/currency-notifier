package currency

import (
	"currency-notifier/config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Currency struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

type Repository struct {
	Config *config.Configuration
}

func Build(config *config.Configuration) *Repository {
	return &Repository{Config: config}
}

func (r *Repository) Get(symbol string) (*Currency, error) {
	response, err := http.Get(fmt.Sprintf("https://api.binance.com/api/v3/ticker/price?symbol=%s", symbol))

	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	c := &Currency{}
	err = json.Unmarshal(data, &c)
	if err != nil {
		return nil, err
	}
	log.Printf("binance response %s: %s", c.Symbol, c.Price)
	return c, nil
}

func (r *Repository) All() ([]*Currency, error) {
	var currencies []*Currency

	for _, currencyConfig := range r.Config.Currencies {
		c, err := r.Get(currencyConfig.Symbol)
		if err != nil {
			return nil, err
		}
		currencies = append(currencies, c)
	}
	return currencies, nil
}
