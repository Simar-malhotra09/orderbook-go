package orderinput

import (
	"Desktop/projects/order-book-simulation/go/orderbook"
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/shopspring/decimal"
)


func ProcessLimitOrder() *orderbook.Order {
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Enter action (A for add, R for remove): ")
    action, _ := reader.ReadString('\n')


    fmt.Print("Enter ticker: ")
    ticker, _ := reader.ReadString('\n')

	fmt.Print("Enter orderbook side: ")
    side, _ := reader.ReadString('\n')

    fmt.Print("Enter price: ")
    priceStr, _ := reader.ReadString('\n')
    price, err := decimal.NewFromString(priceStr)
    if err != nil {
        fmt.Println("Invalid price, please enter a number.")
        return nil
    }

    fmt.Print("Enter size: ")
    sizeStr, _ := reader.ReadString('\n')
    size, err := decimal.NewFromString(sizeStr)
    if err != nil {
        fmt.Println("Invalid size, please enter a number.")
        return nil
    }

    orderID := generateOrderID()
	tickerID := tickerToID()
    timestamp := time.Now()
	side_:= getOrderSide(side)



    // Create a new order
    newOrder := orderbook.NewOrder(orderID, tickerID, side_,orderbook.GoodTillCancel, price, size,size, timestamp)
    return newOrder
}

func getOrderSide(side string)  orderbook.Side{
	var side_ orderbook.Side
	if side == "B" {
		side_ = orderbook.Buy
	} else {
		side_ = orderbook.Sell
	}

	return side_
	
}
func generateOrderID() orderbook.OrderId {
    // Implement a method to generate unique order IDs
    return orderId 
}

func tickerToID() orderbook.TickerId {
  
    return tickerId
}

