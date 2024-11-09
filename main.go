package main

import (
	"Desktop/projects/order-book-simulation/go/orderbook"
	"Desktop/projects/order-book-simulation/go/orderbook_io"

	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
    orderBook := orderbook.NewOrderBook() 

    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Enter action (A for add, R for remove): ")
    action, _ := reader.ReadString('\n')
    action = strings.TrimSpace(action)

    if action != "A" && action != "R" {
        fmt.Println("Invalid action. Only 'A' (add) is supported at the moment.")
        return
    }

    newOrder := orderbook_io.AddLimitOrder() 

    // Print JSON representation of the order
    orderJSON, err := newOrder.MarshalJSON()
    if err != nil {
        log.Fatalf("Error marshaling order: %v", err)
    }
    fmt.Println(string(orderJSON))

    
    side := newOrder.Side() 

	var orderSide *orderbook.OrderSide

	if side == orderbook.Buy {
		orderSide = orderBook.Bids
	} else if side == orderbook.Sell {
		orderSide = orderBook.Asks
	}
	
	orderSide.Append(newOrder)

    
    fmt.Printf("Order added to %s side at price %s\n", side, newOrder.Price().String())
	orderSideJSON,err := orderSide.MarshalJSON()
	if err != nil {
		log.Fatalf("Error marshaling order side to JSON: %v", err)
	}
	
	// Print the JSON representation of the OrderSide
	fmt.Println(string(orderSideJSON))

	orderBookJSON,err:= orderBook.MarshalJSON()
	if err != nil{
		log.Fatalf("Error marshalling orderbook to JSON: %v", err)
	}
	// Print the JSON representation of the OrderSide
	fmt.Println(string(orderBookJSON))
}

