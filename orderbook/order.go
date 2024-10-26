package orderbook

import (
	"encoding/json"
	"fmt"

	"github.com/shopspring/decimal"
)

type Price = decimal.Decimal
type Quantity = decimal.Decimal
type OrderId int64

type Order struct {
    orderId           OrderId
    orderType         OrderType
    side              Side
    price             Price
    initialQuantity   Quantity
    remainingQuantity Quantity
}

func (o *Order) OrderId() OrderId {
    return o.orderId
}

func (o *Order) OrderType() OrderType {
    return o.orderType
}

func (o *Order) Side() Side {
    return o.side
}

func (o *Order) Price() Price {
    return o.price
}

func (o *Order) InitialQuantity() Quantity {
    return o.initialQuantity
}

func (o *Order) RemainingQuantity() Quantity {
    return o.remainingQuantity
}

func (o *Order) FilledQuantity() Quantity {
    initial := decimal.Decimal(o.InitialQuantity())
    remaining := decimal.Decimal(o.RemainingQuantity())
    filled := initial.Sub(remaining)
    return Quantity(filled)
}

func (o *Order) Fill(quantity Quantity) error {
    qty := decimal.Decimal(quantity)
    remaining := decimal.Decimal(o.RemainingQuantity())
    if qty.GreaterThan(remaining) {
        return fmt.Errorf("quantity exceeds remaining quantity for Order ID %d", o.OrderId())
    }
    o.remainingQuantity = Quantity(remaining.Sub(qty))
    return nil
}


func NewOrder(orderId OrderId, side Side, orderType OrderType, price Price, initialQuantity Quantity, remainingQuantity Quantity) *Order {
    return &Order{
        orderId:           orderId,
        side:              side,
        orderType:         orderType,
        price:             price,
        initialQuantity:   initialQuantity,
        remainingQuantity: remainingQuantity,
    }
}


func (o *Order) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		OrderId           OrderId    `json:"orderId"`
		OrderType         OrderType  `json:"orderType"`
		Side              Side       `json:"side"`
		Price             Price      `json:"price"`
		InitialQuantity   Quantity   `json:"initialQuantity"`
		RemainingQuantity Quantity   `json:"remainingQuantity"`
	}{
		OrderId:           o.orderId,
		OrderType:         o.orderType,
		Side:              o.side,
		Price:             o.price,
		InitialQuantity:   o.initialQuantity,
		RemainingQuantity: o.remainingQuantity,
	})
}

func (o *Order) UnmarshalJSON(data []byte) error {
	obj := struct {
		OrderId           OrderId    `json:"orderId"`
		OrderType         OrderType  `json:"orderType"`
		Side              Side       `json:"side"`
		Price             Price      `json:"price"`
		InitialQuantity   Quantity   `json:"initialQuantity"`
		RemainingQuantity Quantity   `json:"remainingQuantity"`
	}{}

	if err := json.Unmarshal(data, &obj); err != nil {
		return err
	}

	o.orderId = obj.OrderId
	o.orderType = obj.OrderType
	o.side = obj.Side
	o.price = obj.Price
	o.initialQuantity = obj.InitialQuantity
	o.remainingQuantity = obj.RemainingQuantity
	return nil
}

