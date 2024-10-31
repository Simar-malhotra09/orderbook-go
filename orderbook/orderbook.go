package orderbook

import (
	"github.com/shopspring/decimal"
)

type OrderBook struct {
	asks map[decimal.Decimal]*OrderQueue 
	bids map[decimal.Decimal]*OrderQueue 

}

func NewOrderBook() *OrderBook {
	return &OrderBook{
		asks: make(map[decimal.Decimal]*OrderQueue),
		bids: make(map[decimal.Decimal]*OrderQueue),
	}
}

