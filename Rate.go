package main

type BuyRate struct {
	client *CoinCheck
}

func (r BuyRate) rate(pair string) string {
	return r.client.Request("GET", "api/rate/"+pair, "")
}
