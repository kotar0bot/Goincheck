package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/BurntSushi/toml"
)

type Balance struct {
	Jpy  string `json:"jpy"`
	Btc  string `json:"btc"`
	Eth  string `json:"eth"`
	Etc  string `json:"etc"`
	Lsk  string `json:"lsk"`
	Fct  string `json:"fct"`
	Xrp  string `json:"xrp"`
	Xem  string `json:"xem"`
	Ltc  string `json:"ltc"`
	Bch  string `json:"bch"`
	Mona string `json:"mona"`
	Xlm  string `json:"xlm"`
	Qtum string `json:"qtum"`
}

type Rate struct {
	Value string `json:"rate"`
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

	var balance Balance
	json.Unmarshal([]byte(client.account.balance()), &balance)

	currencies := []string{"BTC", "ETH", "ETC", "LSK", "FCT", "XRP", "XEM", "LTC", "BCH", "MONA", "XLM", "QTUM"}
	for _, v := range currencies {
		DumpBalance(client, v, GetField(v, balance))
	}
}

func GetField(capital string, balance Balance) string {
	switch capital {
	case "BTC":
		return balance.Btc
	case "ETH":
		return balance.Eth
	case "ETC":
		return balance.Etc
	case "LSK":
		return balance.Lsk
	case "FCT":
		return balance.Fct
	case "XRP":
		return balance.Xrp
	case "XEM":
		return balance.Xem
	case "LTC":
		return balance.Ltc
	case "BCH":
		return balance.Bch
	case "MONA":
		return balance.Mona
	case "XLM":
		return balance.Xlm
	case "QTUM":
		return balance.Qtum
	default:
		return ""
	}
}

func DumpBalance(client CoinCheck, capital string, field string) {
	var rate Rate
	json.Unmarshal([]byte(client.buy_rate.rate(strings.ToLower(capital)+"_jpy")), &rate)
	fmt.Printf("【レート】%s: %s 円", capital, rate.Value)

	fmt.Printf("【所持】%s: %s", capital, field)
	i, _ := strconv.ParseFloat(rate.Value, 64)
	j, _ := strconv.ParseFloat(field, 64)
	fmt.Printf("【所持】%f 円\r\n", i*j)
}
