# Taller 23 - Goroutines Mutex (Mutual Exclusion)

## Información General

**Asignatura:** Programación Orientada a Objetos  
**Taller:** 23 - Goroutines Mutex (Mutual Exclusion)  
**Lenguaje:** Go (Golang)  
**Directorio del proyecto:** `taller23-goroutines`

---

# Propósito

El propósito de este taller es aplicar el uso de **Goroutines**, **WaitGroup** y **Mutex (Mutual Exclusion)** para desarrollar un programa concurrente que procese múltiples órdenes de forma simultánea, evitando condiciones de carrera mediante mecanismos de sincronización.

---

# Objetivo

Desarrollar una aplicación en Go que procese órdenes concurrentemente utilizando tres Goroutines y sincronice el acceso a los datos compartidos mediante Mutex, garantizando la integridad de la información.

---

# Descripción del proyecto

El programa simula el procesamiento de 20 órdenes.

Cada orden posee:

- Un identificador (`ID`)
- Un estado (`Status`)
- Un `Mutex` propio para evitar que varias Goroutines modifiquen la misma orden al mismo tiempo.

Durante la ejecución se crean tres Goroutines que recorren todas las órdenes y actualizan su estado de manera concurrente.

Los posibles estados son:

- Procesando
- Despachando
- Entregado

Además, el programa mantiene un contador global llamado `totalUpdates`, el cual registra el número total de actualizaciones realizadas. Este contador también se protege mediante un Mutex independiente para evitar condiciones de carrera.

---

# Tecnologías utilizadas

- Go (Golang)
- Goroutines
- sync.WaitGroup
- sync.Mutex
- Visual Studio Code
- Git
- GitHub

---

# Estructura del proyecto

```text
taller23-goroutines/
│
├── README.md
└── src/
    ├── go.mod
    └── main.go
```

---

# Instrucciones de ejecución

1. Abrir una terminal.
2. Ingresar al directorio del proyecto.

```bash
cd taller23-goroutines/src
```

3. Ejecutar el programa.

```bash
go run .
```

---

# Desarrollo del taller

## Iteración 1

### Actividades realizadas

- Creación del proyecto.
- Importación de librerías.
- Declaración de la estructura `Order`.
- Declaración de variables globales.
- Creación del Mutex para sincronización.

### Código implementado

- Struct `Order`
- Variable `totalUpdates`
- Variable `updateMutex`

### Evidencia de ejecución

```text
Taller 23 - Goroutines y Mutex
```

### Resultado

Se verificó que la estructura base del proyecto compiló correctamente.

---

## Iteración 2

### Actividades realizadas

- Implementación de la función `generateOrders()`.
- Generación de 20 órdenes.
- Creación de tres Goroutines.
- Uso de `sync.WaitGroup`.
- Procesamiento concurrente de las órdenes.

### Evidencia de ejecución

```text
Orden 1
Orden 2
Orden 3
...
Orden 20

Todas las operaciones completadas.
Total Actualizaciones 0
```

### Resultado

Se comprobó el funcionamiento de las Goroutines ejecutando múltiples procesos simultáneamente.

En esta etapa todavía no se protegían los recursos compartidos mediante Mutex.

---

## Iteración 3

### Actividades realizadas

- Implementación de Mutex para cada orden.
- Protección del contador global mediante `updateMutex`.
- Actualización segura del estado de las órdenes.
- Reporte final de estados.

### Evidencia de ejecución

```text
Actualizando orden 1 con estado: Despachando
Actualizando orden 1 con estado: Procesando
Actualizando orden 2 con estado: Procesando
Actualizando orden 3 con estado: Entregado
...
Actualizando orden 20 con estado: Entregado

Estado final de órdenes:

Orden 1 -> Entregado
Orden 2 -> Procesando
Orden 3 -> Despachando
Orden 4 -> Procesando
Orden 5 -> Procesando
Orden 6 -> Procesando
Orden 7 -> Entregado
Orden 8 -> Procesando
Orden 9 -> Despachando
Orden 10 -> Despachando
Orden 11 -> Entregado
Orden 12 -> Despachando
Orden 13 -> Despachando
Orden 14 -> Entregado
Orden 15 -> Entregado
Orden 16 -> Procesando
Orden 17 -> Despachando
Orden 18 -> Entregado
Orden 19 -> Procesando
Orden 20 -> Entregado

Todas las operaciones completadas.

Total Actualizaciones 60
```

### Resultado

La ejecución confirmó que:

- Cada orden fue actualizada correctamente.
- Las tres Goroutines trabajaron de forma concurrente.
- El uso de Mutex evitó condiciones de carrera.
- El contador global registró correctamente las 60 actualizaciones realizadas.

---

# Conceptos aplicados

## Goroutines

Las Goroutines permiten ejecutar funciones concurrentemente, facilitando la realización de múltiples tareas al mismo tiempo.

## WaitGroup

El `sync.WaitGroup` permite esperar a que todas las Goroutines finalicen antes de terminar la ejecución del programa.

## Mutex

El `sync.Mutex` garantiza la exclusión mutua, evitando que dos Goroutines accedan simultáneamente a un mismo recurso compartido.

En este proyecto se utilizaron dos tipos de Mutex:

- Un Mutex para proteger cada orden.
- Un Mutex para proteger el contador global de actualizaciones.

---

# Resultados obtenidos

Durante el desarrollo del taller se logró:

- Crear una aplicación concurrente utilizando Goroutines.
- Sincronizar la ejecución mediante WaitGroup.
- Evitar condiciones de carrera mediante Mutex.
- Actualizar correctamente las 20 órdenes.
- Obtener un total de 60 actualizaciones registradas.
- Mostrar el estado final de todas las órdenes procesadas.

---

# Conclusiones

- Las Goroutines permiten ejecutar múltiples tareas de forma concurrente, optimizando el procesamiento de información.
- El uso de WaitGroup garantiza que el programa espere correctamente la finalización de todos los procesos concurrentes.
- Mutex es una herramienta fundamental para proteger recursos compartidos y evitar condiciones de carrera.
- La combinación de Goroutines, WaitGroup y Mutex permite desarrollar aplicaciones concurrentes seguras, eficientes y organizadas en Go.
- Este taller permitió comprender la importancia de la sincronización cuando varios procesos acceden simultáneamente a los mismos datos.