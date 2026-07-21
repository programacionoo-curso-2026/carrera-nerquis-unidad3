# Taller 20 - Goroutines en Go

## Objetivo

Practicar el uso de goroutines en Go para ejecutar tareas concurrentes.

## Descripción

El programa crea 10 goroutines utilizando la palabra reservada go.

Cada goroutine ejecuta la función ShowGoroutine y muestra su identificador junto con un tiempo de espera aleatorio.

## Ejecución

Ejecutar el programa con:

go run src/main.go

## Decisiones de diseño

- Se utilizó time.Sleep para controlar la duración del programa.
- Se utilizó rand para generar tiempos aleatorios.
- Se utilizó fmt.Printf para mostrar información en consola.

## Resultado esperado

El programa muestra las 10 goroutines ejecutándose en un orden diferente debido a la concurrencia.