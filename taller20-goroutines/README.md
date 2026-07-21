# Taller 20 - Goroutines en Go

## Objetivo

Practicar el uso de goroutines en Go para ejecutar tareas concurrentes y comprender el funcionamiento de la programación paralela.

## Descripción

El programa crea 10 goroutines utilizando la palabra reservada `go`.

Cada goroutine ejecuta la función `ShowGoroutine`, recibe un identificador y muestra un tiempo de espera aleatorio antes de finalizar.

Debido a la ejecución concurrente, las goroutines pueden terminar en diferente orden en cada ejecución.

## Estructura del proyecto

```text
taller20-goroutines
│
├── README.md
│
└── src
    └── main.go
