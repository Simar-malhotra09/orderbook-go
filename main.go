package main

import (
	"Desktop/projects/order-book-simulation/go/orderbook"
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/shopspring/decimal"
)

func main() {

	// Create a new OrderSide for asks and bids in the order book
	asks := orderbook.NewOrderSide()
	bids := orderbook.NewOrderSide()
	_ = bids

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter trades  or 'exit' to quit:")

	for scanner.Scan() {
		line := scanner.Text()
		if line == "exit" {
			break
		}

		fields := strings.Fields(line)
		// first two fields will always be the action and orderId.
		// actions are referred by their assigned string literals:
		// A -> Add order
		// M -> Modify order
		// C -> Cancel order

		action := fields[0]
		orderId, err := strconv.ParseInt(fields[1], 10, 64)  // Convert orderId to int64
		if err != nil {
			log.Fatalf("Failed to parse orderId: %v", err)
		}
		

		if action == "A" {
			// To add order, first figure out which side
			// of the orderbook to add to
			var side orderbook.Side
			if fields[2] == "B" {
				side = orderbook.Buy
			} else {
				side = orderbook.Sell
			}
		    
			
			price, err := decimal.NewFromString(fields[3])
			if err != nil {
				log.Fatalf("Failed to parse price: %v", err)
			}
		
			quantity, err := decimal.NewFromString(fields[4])
			if err != nil {
				log.Fatalf("Failed to parse quantity: %v", err)
			}
		
			// Create the new order
			// Order type is arbitary right now
			newOrder := orderbook.NewOrder(orderbook.OrderId(orderId), side, orderbook.GoodTillCancel, price, quantity, quantity)
			fmt.Printf("Created new order:")


			// Print the JSON representation
			orderJSON, err := json.Marshal(newOrder)
			if err != nil {
				log.Fatalf("Error marshaling order to JSON: %v", err)
			}
			fmt.Println("Order JSON:", string(orderJSON))


            // add order to correct side
			
			if side == orderbook.Buy{
				bids.Append(newOrder)
				//  Marshal the order to JSON
				bidsJSON, err := json.Marshal(bids)
				if err != nil {
					log.Fatalf("Error marshaling order to JSON: %v", err)
				}
	
				// Print the JSON representation
				fmt.Println("bids JSON:", string(bidsJSON))

			}else{
				asks.Append(newOrder)
				//  Marshal the order to JSON
				asksJSON, err := json.Marshal(asks)
				if err != nil {
					log.Fatalf("Error marshaling order to JSON: %v", err)
				}
	
				// Print the JSON representation
				fmt.Println("asks JSON:", string(asksJSON))
				
			}







		} else {
			continue
		}



		
		fmt.Println("Received input:", line)
	}
}
