package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Order struct {
	ID     int
	Status string
	mu     sync.Mutex
}

var (
	totalUpdates int
	updateMutex  sync.Mutex
)

func generateOrders(n int) []*Order {
	orders := make([]*Order, n)

	for i := 0; i < n; i++ {
		orders[i] = &Order{
			ID:     i + 1,
			Status: "Pendiente",
		}
	}

	return orders
}

func updateOrderStatus(order *Order) {
	fmt.Printf("Orden %d\n", order.ID)
}

func main() {

	rand.Seed(time.Now().UnixNano())

	var wg sync.WaitGroup

	wg.Add(3)

	orders := generateOrders(20)

	for i := 0; i < 3; i++ {

		go func() {

			defer wg.Done()

			for _, order := range orders {
				updateOrderStatus(order)
			}

		}()

	}

	wg.Wait()

	fmt.Println("Todas las operaciones completadas.")
	fmt.Printf("Total Actualizaciones %d\n", totalUpdates)

}
