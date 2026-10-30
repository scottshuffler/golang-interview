package main

import "fmt"

type Order struct {
	Customer string
	Amount   int
}

func totalsByCustomer(orders []Order) (map[string]int, error) {
	totals := make(map[string]int)

	for _, order := range orders {
		totals[order.Customer] = order.Amount
	}

	return totals, nil
}

func main() {
	orders := []Order{
		{Customer: "Alice", Amount: 10},
		{Customer: "", Amount: 2},
		{Customer: "Bob", Amount: 7},
		{Customer: "Alice", Amount: 5},
		{Customer: "", Amount: 20},
	}

	totals, _ := totalsByCustomer(orders)

	fmt.Println("Alice:", totals["Alice"])
	fmt.Println("Bob:", totals["Bob"])
}
