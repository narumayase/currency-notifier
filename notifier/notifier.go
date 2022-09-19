package notifier

import (
	"currency-notifier/config"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Notifier struct {
	Config *config.Configuration
}

func Build(config *config.Configuration) *Notifier {
	return &Notifier{Config: config}
}

func (n *Notifier) Send(msg string) error {
	response, err := http.Get(fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%s&text=%s", n.Config.TelegramData.Token, n.Config.TelegramData.ChatId, msg))

	if err != nil {
		return err
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	log.Printf("telegram response: %s", string(data))
	return nil
}
