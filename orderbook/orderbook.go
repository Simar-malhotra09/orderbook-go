package orderbook

import (
	"container/list"
)

type OrderBook struct{
	orders  map[string]*list.Element
	asks *OrderSide
	bids *OrderSide
}