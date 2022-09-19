package main

import (
	"currency-notifier/config"
	"currency-notifier/currency"
	"currency-notifier/notifier"
	"fmt"
	"log"
	"strconv"
)

func main() {
	configuration := config.Load()
	repo := currency.Build(configuration)
	n := notifier.Build(configuration)

	currencies, err := repo.All()
	if err != nil {
		log.Fatalf("there was an error getting currencies info %s", err.Error())
	}
	for _, c := range currencies {
		p := GetFloatPrice(c.Price)
		if p > configuration.GetLimMax(c.Symbol) || p < configuration.GetLimMin(c.Symbol) {
			err := n.Send(fmt.Sprintf("%s: %s", c.Symbol, c.Price))
			if err != nil {
				log.Fatalf("there was an error getting currencies info %s", err.Error())
			}
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
