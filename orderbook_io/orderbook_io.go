package orderbook_io

import (
	"Desktop/projects/order-book-simulation/go/orderbook"
	"Desktop/projects/order-book-simulation/go/tickers"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gofrs/uuid/v5"
	"github.com/shopspring/decimal"
)

func AddLimitOrder() *orderbook.Order {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter ticker: ")
	ticker, _ := reader.ReadString('\n')
	ticker = strings.TrimSpace(ticker)

	tickerID, err := tickers.TickerToID(ticker)
	if err != nil {
		fmt.Println("Invalid ticker symbol:", err)
		return nil
	}

	fmt.Print("Enter orderbook side (B for buy, S for sell): ")
	side, _ := reader.ReadString('\n')
	side = strings.TrimSpace(side)
	orderSide := GetOrderSide(side)


	fmt.Print("Enter price: ")
	priceStr, _ := reader.ReadString('\n')
	priceStr = strings.TrimSpace(priceStr)
	price, err := decimal.NewFromString(priceStr)
	if err != nil {
		fmt.Println("Invalid price. Please enter a number.")
		return nil
	}

	fmt.Print("Enter size: ")
	sizeStr, _ := reader.ReadString('\n')
	sizeStr = strings.TrimSpace(sizeStr)
	size, err := decimal.NewFromString(sizeStr)
	if err != nil {
		fmt.Println("Invalid size. Please enter a number.")
		return nil
	}

	orderID := GenerateOrderID()
	timestamp := time.Now()

	
	newOrder := orderbook.NewOrder(orderID, tickerID, orderSide, orderbook.GoodTillCancel, price, size, size, timestamp)
		

    return newOrder
}

func GetOrderSide(side string) orderbook.Side {
	switch strings.ToUpper(side) {
	case "B":
		return orderbook.Buy
	case "S":
		return orderbook.Sell
	default:
		return orderbook.Buy
	}
}

func GenerateOrderID() orderbook.OrderId {
	orderID, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("Failed to generate UUID: %v", err)
	}
	log.Printf("Generated Order ID: %v", orderID)
	return orderbook.OrderId(orderID.String())
}

