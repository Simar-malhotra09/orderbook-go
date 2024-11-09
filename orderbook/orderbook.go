package orderbook

import (
	"container/list"
	"encoding/json"
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


// MarshalJSON implements json.Marshaler interface for OrderBook
func (ob *OrderBook) MarshalJSON() ([]byte, error) {
	
	return json.Marshal(&struct {

		Asks   *OrderSide              `json:"asks"`
		Bids   *OrderSide              `json:"bids"`
	}{

		Asks:   ob.Asks,
		Bids:   ob.Bids,
	})
}


func (ob *OrderBook) UnmarshalJSON(data []byte) error {
	// Temporary struct to hold unmarshalled JSON data
	var temp struct {

		Asks   *OrderSide              `json:"asks"`
		Bids   *OrderSide              `json:"bids"`
	}


	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	

	ob.Asks = temp.Asks
	ob.Bids = temp.Bids

	return nil
}
