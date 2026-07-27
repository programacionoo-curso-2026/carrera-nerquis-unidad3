# Taller 22 - GoRoutines

## Objetivo

El objetivo de este taller es aplicar el uso de **GoRoutines** en el lenguaje Go para ejecutar procesos concurrentes y comprender el manejo de tareas simultáneas mediante goroutines y sincronización con `sync.WaitGroup`.

## Descripción del proyecto

Este proyecto simula un sistema de procesamiento de órdenes. Cada orden posee un identificador y un estado que cambia durante la ejecución del programa.

El sistema realiza las siguientes operaciones:

- Generación de órdenes.
- Procesamiento de órdenes.
- Actualización de estados.
- Reporte periódico del estado de las órdenes.
- Ejecución concurrente mediante goroutines.

Para controlar la finalización de las tareas concurrentes se utiliza `sync.WaitGroup`.

## Estructura del proyecto

```
taller22-goroutines
│
├── README.md
│
└── src
    └── main.go
```

## Tecnologías utilizadas

- Lenguaje Go
- GoRoutines
- Sync WaitGroup
- Git y GitHub para control de versiones

## Implementación por iteraciones

### Iteración 1: Estructura inicial

Se creó la estructura del proyecto y la definición del tipo `Order`.

Se implementó:

- Struct `Order`.
- Función `generateOrders()`.
- Creación de 20 órdenes iniciales.

---

### Iteración 2: Procesamiento de órdenes

Se agregó la función `processOrders()` para simular el procesamiento de cada orden utilizando tiempos aleatorios.

---

### Iteración 3: Actualización de estados

Se implementó la función `updateOrderStatuses()` para cambiar los estados de las órdenes.

Estados utilizados:

- Procesando.
- Despachando.
- Entregado.

---

### Iteración 4: Reporte de estados

Se creó la función `reportOrderStatus()` encargada de mostrar periódicamente el estado actual de todas las órdenes.

---

### Iteración 5: Ejecución secuencial

Se probó la ejecución de las funciones principales de forma secuencial para verificar el funcionamiento inicial.

---

### Iteración 6: Implementación de GoRoutines

Se incorporó el uso de la palabra reservada `go` para ejecutar tareas concurrentes:

- Procesamiento de órdenes.
- Actualización de estados.
- Generación de reportes.

---

### Iteración 7: Sincronización

Se agregó `sync.WaitGroup` para controlar la finalización correcta de las goroutines.

---

### Iteración 8: Versión final

Se implementaron funciones anónimas con:

- `wg.Add()`
- `wg.Done()`
- `wg.Wait()`

permitiendo una ejecución concurrente controlada.

## Ejecución del programa

Desde la raíz del repositorio ejecutar:

```bash
go run taller22-goroutines/src/main.go
```

## Resultado esperado

Durante la ejecución se muestran mensajes similares a:

```
Procesando orden 1
Actualizando orden 2 con estado: Entregado

--- Reporte Estado de las Ordenes ---
Orden 1: Procesando
Orden 2: Entregado
---------------------------------------

Todas las operaciones completadas. Saliendo
```

El orden de los mensajes puede variar debido a la ejecución concurrente de las goroutines.

## Control de versiones

El desarrollo del taller fue registrado mediante commits progresivos utilizando Git.

Ejemplos de commits:

- `taller22: crea estructura inicial de goroutines`
- `taller22: implementa generacion y procesamiento de ordenes`

Cada avance fue documentado para evidenciar el proceso de desarrollo.

## Autor

**Nerquis Carrera**

Ingeniería en Sistemas de la Información  
Universidad Internacional del Ecuador