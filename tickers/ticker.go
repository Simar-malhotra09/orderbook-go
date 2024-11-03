package tickers

import (
	"Desktop/projects/order-book-simulation/go/orderbook"
	"errors"
)



type Ticker struct {
	ID      orderbook.TickerId
	Symbol   string
	Name     string
	Exchange string
}


func TickerToID(symbol string) (orderbook.TickerId, error) {
	ticker, found := TickerLookup[symbol]
	if !found {
		return 0, errors.New("ticker symbol not found")
	}
	return ticker.ID, nil
}
var TickerLookup = map[string]Ticker{
	"AAPL":  {ID: 1, Symbol: "AAPL", Name: "Apple Inc.", Exchange: "NASDAQ"},
	"GOOGL": {ID: 2, Symbol: "GOOGL", Name: "Alphabet Inc.", Exchange: "NASDAQ"},
	"AMZN":  {ID: 3, Symbol: "AMZN", Name: "Amazon.com Inc.", Exchange: "NASDAQ"},
	"MSFT":  {ID: 4, Symbol: "MSFT", Name: "Microsoft Corporation", Exchange: "NASDAQ"},
	"TSLA":  {ID: 5, Symbol: "TSLA", Name: "Tesla Inc.", Exchange: "NASDAQ"},
	"FB":    {ID: 6, Symbol: "FB", Name: "Meta Platforms, Inc.", Exchange: "NASDAQ"},
	"NVDA":  {ID: 7, Symbol: "NVDA", Name: "NVIDIA Corporation", Exchange: "NASDAQ"},
	"BRK.B": {ID: 8, Symbol: "BRK.B", Name: "Berkshire Hathaway Inc.", Exchange: "NYSE"},
	"JPM":   {ID: 9, Symbol: "JPM", Name: "JPMorgan Chase & Co.", Exchange: "NYSE"},
	"V":     {ID: 10, Symbol: "V", Name: "Visa Inc.", Exchange: "NYSE"},
	"JNJ":   {ID: 11, Symbol: "JNJ", Name: "Johnson & Johnson", Exchange: "NYSE"},
	"WMT":   {ID: 12, Symbol: "WMT", Name: "Walmart Inc.", Exchange: "NYSE"},
	"PG":    {ID: 13, Symbol: "PG", Name: "Procter & Gamble Co.", Exchange: "NYSE"},
	"UNH":   {ID: 14, Symbol: "UNH", Name: "UnitedHealth Group Incorporated", Exchange: "NYSE"},
	"DIS":   {ID: 15, Symbol: "DIS", Name: "The Walt Disney Company", Exchange: "NYSE"},
	"HD":    {ID: 16, Symbol: "HD", Name: "The Home Depot, Inc.", Exchange: "NYSE"},
	"VZ":    {ID: 17, Symbol: "VZ", Name: "Verizon Communications Inc.", Exchange: "NYSE"},
	"INTC":  {ID: 18, Symbol: "INTC", Name: "Intel Corporation", Exchange: "NASDAQ"},
	"CMCSA": {ID: 19, Symbol: "CMCSA", Name: "Comcast Corporation", Exchange: "NASDAQ"},
	"NKE":   {ID: 20, Symbol: "NKE", Name: "Nike, Inc.", Exchange: "NYSE"},
}
