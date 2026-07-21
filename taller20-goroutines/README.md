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
```

## Ejecución

Para ejecutar el programa, ubicarse dentro de la carpeta del taller y ejecutar el siguiente comando:

```bash
go run src/main.go
```

## Decisiones de diseño

- Se utilizó la palabra reservada `go` para crear goroutines y ejecutar tareas concurrentes.
- Se utilizó `math/rand` para generar tiempos de espera aleatorios.
- Se utilizó `fmt.Printf` para mostrar información de cada goroutine en consola.
- Se utilizó `sync.WaitGroup` para controlar la finalización correcta de todas las goroutines.
- Se agregaron comentarios dentro del código para facilitar la comprensión del funcionamiento del programa.

## Resultado esperado

El programa muestra las 10 goroutines ejecutándose de manera concurrente.

Debido a la ejecución paralela, el orden de finalización puede variar en cada ejecución.

Ejemplo de salida:

```text
Goroutine #3 con retraso de 242 ms
Goroutine #1 con retraso de 99 ms
Goroutine #4 con retraso de 447 ms
Todas las goroutines finalizaron correctamente
```

## Conclusión

En este taller se practicó el uso de goroutines en Go, comprendiendo cómo ejecutar tareas concurrentes y cómo el orden de finalización puede variar debido a la ejecución paralela.

El uso de goroutines permite desarrollar programas capaces de realizar múltiples tareas al mismo tiempo y aprovechar mejor los recursos del sistema.
