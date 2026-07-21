package main

import (
	"fmt"
	"math/rand"
	"time"
)

func ShowGoroutine(id int) {
	delay := rand.Intn(500)

	fmt.Printf("Goroutine #%d con retraso de %d ms\n", id, delay)

	time.Sleep(time.Duration(delay) * time.Millisecond)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 10; i++ {
		go ShowGoroutine(i + 1)
	}

	time.Sleep(1 * time.Second)
}
