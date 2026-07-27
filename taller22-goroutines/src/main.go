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
}

func main() {

	var wg sync.WaitGroup
	wg.Add(3)

	orders := generateOrders(20)

	go func() {
		defer wg.Done()
		processOrders(orders)
	}()

	go func() {
		defer wg.Done()
		updateOrderStatuses(orders)
	}()

	go func() {
		defer wg.Done()
		reportOrderStatus(orders)
	}()

	wg.Wait()

	fmt.Println("Todas las operaciones completadas. Saliendo")
}

func generateOrders(count int) []*Order {

	orders := make([]*Order, count)

	for i := 0; i < count; i++ {
		orders[i] = &Order{
			ID:     i + 1,
			Status: "Pendiente",
		}
	}

	return orders
}

func processOrders(orders []*Order) {

	for _, order := range orders {

		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)

		fmt.Printf("Procesando orden %d\n", order.ID)

		order.Status = "Procesando"
	}
}

func updateOrderStatuses(orders []*Order) {

	statuses := []string{
		"Procesando",
		"Despachando",
		"Entregado",
	}

	for _, order := range orders {

		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)

		status := statuses[rand.Intn(len(statuses))]

		order.Status = status

		fmt.Printf(
			"Actualizando orden %d con estado: %s\n",
			order.ID,
			status,
		)
	}
}

func reportOrderStatus(orders []*Order) {

	for i := 0; i < 5; i++ {

		time.Sleep(1 * time.Second)

		fmt.Println("\n--- Reporte Estado de las Ordenes ---")

		for _, order := range orders {

			fmt.Printf(
				"Orden %d: %s\n",
				order.ID,
				order.Status,
			)
		}

		fmt.Println("---------------------------------------")
	}
}
