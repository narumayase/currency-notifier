# currency-notifier

Busca en la url de binance las cryptocurrencies configuradas en el archivo [config.json](https://github.com/narumayase/currency-notifier/blob/main/conf.json) con sus valores mínimos y máximos para notificar a un bot de Telegram.

## Empezando 🚀

### Configuración

Configurar los datos del bot de telegram:

```
"telegramData": {
    "token": "5649405085:AAEGtL1PMSNbOLLRUq0tiGoZVgCs",
    "chatId": "3601488"
  },
```

en el archivo [config.json](https://github.com/narumayase/currency-notifier/blob/main/conf.json) con el siguiente formato:

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

### Requerimientos

- Go go1.18.4+.

### Ejemplo ejecutando localmente

* Run:

```
$ go build
$ go run main.go
```

