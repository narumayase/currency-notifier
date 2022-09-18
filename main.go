package main

import (
	"currency-notifier/config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func main() {
	configuration := config.Load()
	currencies := GetCurrencyInfo(configuration)

	for _, c := range currencies {
		p := GetFloatPrice(c.Price)
		if p > configuration.GetLimMax(c.Symbol) || p < configuration.GetLimMin(c.Symbol) {
			log.Printf("%s: %s", c.Symbol, c.Price)
			Notify(configuration, fmt.Sprintf("%s: %s", c.Symbol, c.Price))
		}
	}
}

func GetFloatPrice(price string) float64 {
	if p, err := strconv.ParseFloat(price, 64); err == nil {
		return p
	} else {
		log.Fatal(err)
	}
	return 0
}

func GetCurrencyInfo(config *config.Configuration) []*Currency {
	var currencies []*Currency

	for _, currencyConfig := range config.Currencies {
		response, err := http.Get(fmt.Sprintf("https://api.binance.com/api/v3/ticker/price?symbol=%s", currencyConfig.Symbol))

		if err != nil {
			log.Fatal(err)
		}
		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		c := &Currency{}
		err = json.Unmarshal(data, &c)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("binance response %s: %s", c.Symbol, c.Price)

		currencies = append(currencies, c)
	}
	return currencies
}

func Notify(config *config.Configuration, msg string) {
	response, err := http.Get(fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%s&text=%s", config.TelegramData.Token, config.TelegramData.ChatId, msg))

	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("telegram response: %s", string(data))
}

type Currency struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}
