# currency-notifier

Finds in the binance url the cryptocurrencies configured in the file [config.json](https://github.com/narumayase/currency-notifier/blob/main/conf.json) and then sends a notification to a Telegram bot if it exeeds the max or min limits of the currencies configured.

## Starting 🚀

### Configuration

Configure Telegram bot:

```
"telegramData": {
    "token": "5649405085:AAEGtL1PMSNbOLLRUq0tiGoZVgCs",
    "chatId": "3601488"
  },
```

in the file [config.json](https://github.com/narumayase/currency-notifier/blob/main/conf.json) with format:

```
{
  "telegramData": {
    "token": "5649405085:AAEGtL1PMSNbOLLRUq0tiGoZVgCs",
    "chatId": "3601488"
  },
  "currencies": [
    {
      "symbol": "BTCUSDT",
      "limMax": 20029.82000000,
      "limMin": 18000.00000000
    },
    {
      "symbol": "ETHUSDT",
      "limMax": 1600.00000000,
      "limMin": 1100.00000000
    },
    {
      "symbol": "ETCUSDT",
      "limMax": 1600.00000000,
      "limMin": 1100.00000000
    }
  ]
}
```

and then you can configure the limMin and limMax.
Looking at the showed example, if BTUSDT exceeds limMax value (20029.82000000) or if BTCUSDT is below limMin (18000.00000000) it will send a telegram notification.

### Requirements

- Go go1.18.4+.

### Example running locally

* Run:

```
$ go build
$ go run main.go
```
