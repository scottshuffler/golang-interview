package main

import "fmt"

type Order struct {
	Customer string
	Amount   int
}

func totalsByCustomer(orders []Order) (map[string]int, error) {
	totals := make(map[string]int)
	var firstErr error

	for _, order := range orders {
		if order.Customer == "" || order.Amount <= 0 {
			if firstErr == nil {
				firstErr = fmt.Errorf("invalid order for customer %q", order.Customer)
			}
			continue
		}

		totals[order.Customer] = order.Amount
	}

	return totals, firstErr
}

func main() {
	orders := []Order{
		{Customer: "Alice", Amount: 10},
		{Customer: "Bob", Amount: 7},
		{Customer: "Alice", Amount: 5},
		{Customer: "", Amount: 20},
	}

	totals, err := totalsByCustomer(orders)
	if err != nil {
		fmt.Println("warning:", err)
	}

	fmt.Println("Alice:", totals["Alice"])
	fmt.Println("Bob:", totals["Bob"])
}
