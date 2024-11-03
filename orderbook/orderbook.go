package orderbook

import (
	"container/list"
)

type OrderBook struct {
	orders map[string]*list.Element // orderID -> *Order (*list.Element.Value.(*Order))

	Asks *OrderSide
	Bids *OrderSide
}


func NewOrderBook() *OrderBook {
	return &OrderBook{
		orders: map[string]*list.Element{},
		Bids:   NewOrderSide(),
		Asks:   NewOrderSide(),
	}
}


