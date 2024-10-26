package main

import (
	"fmt"

	"Desktop/projects/order-book-simulation/go/orderbook"

	"github.com/shopspring/decimal"
)


func main() {
	// Initialize an OrderQueue with a specific price level
	price := decimal.NewFromFloat(100.0)
	orderQueue := orderbook.NewOrderQueue(price)

	// Create initial orders and append them to the queue
	order1 := orderbook.NewOrder(1, orderbook.Buy, orderbook.GoodTillCancel, price, decimal.NewFromFloat(5.0), decimal.NewFromFloat(5.0))
	order2 := orderbook.NewOrder(2, orderbook.Sell, orderbook.FillAndKill, price, decimal.NewFromFloat(3.0), decimal.NewFromFloat(3.0))

	orderQueue.Append(order1)
	element := orderQueue.Append(order2) // Keep reference to this element to update it

	// Print initial state of the queue
	fmt.Println("Initial Order Queue:")
	fmt.Println(orderQueue)

	// Create a new order to replace order2
	newOrder := orderbook.NewOrder(2, orderbook.Sell, orderbook.FillAndKill, price, decimal.NewFromFloat(4.0), decimal.NewFromFloat(4.0))

	// Use Update to replace order2 with the new order
	orderQueue.Update(element, newOrder)

	// Print updated state of the queue
	fmt.Println("Updated Order Queue:")
	fmt.Println(orderQueue)
}
// func main() {
//     // Create an example order
//     order := orderbook.NewOrder(
//         1, // Order ID
//         orderbook.Buy, // Side (using the Buy constant)
//         orderbook.GoodTillCancel, // Order Type (using the GoodTillCancel constant)
//         decimal.NewFromFloat(100.50), // Price
//         decimal.NewFromInt(10), // Initial Quantity
//         decimal.NewFromInt(10), // Remaining Quantity
//     )

//     // Display the order details
//     fmt.Printf("Order ID: %d\n", order.OrderId())
//     fmt.Printf("Order Type: %s\n", order.OrderType())
//     fmt.Printf("Side: %s\n", order.Side())
//     fmt.Printf("Price: %s\n", order.Price().String())
//     fmt.Printf("Initial Quantity: %s\n", order.InitialQuantity().String())
//     fmt.Printf("Remaining Quantity: %s\n", order.RemainingQuantity().String())
// }