package main

import (
	"fmt"
	"math/rand"
	"time"
)

// ShowGoroutine simula la ejecución de una goroutine.
// Recibe un identificador y espera un tiempo aleatorio.
func ShowGoroutine(id int) {

	// Genera un retraso aleatorio entre 0 y 499 milisegundos.
	delay := rand.Intn(500)

	fmt.Printf("Goroutine #%d con retraso de %d ms\n", id, delay)

	// Simula una tarea que tarda cierto tiempo en finalizar.
	time.Sleep(time.Duration(delay) * time.Millisecond)
}

func main() {

	// Inicializa la semilla para obtener valores aleatorios diferentes.
	rand.Seed(time.Now().UnixNano())

	// Crea 10 goroutines ejecutándose de manera concurrente.
	for i := 0; i < 10; i++ {
		go ShowGoroutine(i + 1)
	}

	// Espera para permitir que todas las goroutines terminen.
	time.Sleep(1 * time.Second)
}
