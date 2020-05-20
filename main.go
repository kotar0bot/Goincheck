package main

import (
	"encoding/json"

	"github.com/BurntSushi/toml"
)

type Balance struct {
	Jpy string `json:"jpy"`
}

type Config struct {
	APIKey string `toml:"apikey"`
	Secret string `toml:"secret"`
}

func main() {
	var config Config
	_, err := toml.DecodeFile("config.tml", &config)
	if err != nil {
		panic(err)
	}
	client := new(CoinCheck).NewClient(config.APIKey, config.Secret)

	// /** Public API */
	// client.ticker.all()
	// client.trade.all()
	// client.order_book.all()

	// // 取引履歴
	// client.order.transactions()

	// 残高
	var balance Balance
	json.Unmarshal([]byte(client.account.balance()), &balance)
	println("=============================================")
	println(balance.Jpy)
}
