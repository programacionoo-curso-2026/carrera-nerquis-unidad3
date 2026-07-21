package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

// ShowGoroutine simula la ejecución de una goroutine.
// Recibe un identificador y espera un tiempo aleatorio.
func ShowGoroutine(id int) {

	defer wg.Done()

	// Genera un retraso aleatorio entre 0 y 499 milisegundos.
	delay := rand.Intn(500)

	fmt.Printf("Goroutine #%d con retraso de %d ms\n", id, delay)

	// Simula una tarea que tarda cierto tiempo en finalizar.
	time.Sleep(time.Duration(delay) * time.Millisecond)
}

func main() {

	// Inicializa valores aleatorios diferentes en cada ejecución.
	rand.Seed(time.Now().UnixNano())

	// Crea 10 goroutines ejecutándose de manera concurrente.
	for i := 0; i < 10; i++ {

		wg.Add(1)

		go ShowGoroutine(i + 1)
	}

	// Espera hasta que todas las goroutines terminen.
	wg.Wait()

	fmt.Println("Todas las goroutines finalizaron correctamente")
}