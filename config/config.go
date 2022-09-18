package config

import (
	"encoding/json"
	"log"
	"os"
)

type Configuration struct {
	TelegramData *TelegramData
	Currencies   []*Currency
}

type TelegramData struct {
	Token  string
	ChatId string
}

type Currency struct {
	Symbol string
	LimMax float64
	LimMin float64
}

func (c *Configuration) Json() string {
	s, err := json.Marshal(c)
	if err != nil {
		return ""
	}
	return string(s)
}

func Load() *Configuration {
	file, err := os.Open("conf.json")
	if err != nil {
		log.Fatalf("conf.json is missing")
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	conf := Configuration{}
	err = decoder.Decode(&conf)
	if err != nil {
		log.Fatal("error decoding json file:", err)
	}
	log.Printf("configuration loaded: %s\n", conf.Json())
	return &conf
}

func (c *Configuration) GetLimMax(name string) float64 {
	for _, currency := range c.Currencies {
		if currency.Symbol == name {
			return currency.LimMax
		}
	}
	return 0
}

func (c *Configuration) GetLimMin(name string) float64 {
	for _, currency := range c.Currencies {
		if currency.Symbol == name {
			return currency.LimMin
		}
	}
	return 0
}
