# Bitácora de desarrollo - Unidad 3

## 20/07/2026 - Taller 20: Goroutines - Fase 1: Creación del proyecto

### Qué hice:
- Se creó el repositorio del proyecto en GitHub dentro de la organización programacionoo-curso-2026.
- Se clonó el repositorio en la computadora.
- Se creó la carpeta correspondiente al Taller 20.
- Se organizó la estructura del proyecto:
  - taller20-goroutines
  - src
  - main.go
- Se implementó el programa utilizando goroutines en Go.

### Qué problema encontré:
- El repositorio inicialmente fue clonado dentro de una carpeta incorrecta, generando una estructura duplicada.
- Al ejecutar el programa se identificaron errores de sintaxis en el código proporcionado originalmente.

### Cómo lo resolví:
- Se eliminó la copia incorrecta del repositorio y se realizó nuevamente la clonación desde la ubicación correcta.
- Se corrigió el código para utilizar correctamente los paquetes de Go, la función de impresión y el manejo del tiempo de espera.

### Próximo paso:
- Ejecutar el programa y verificar el funcionamiento correcto de las goroutines.


---

## 20/07/2026 - Taller 20: Goroutines - Fase 2: Implementación y pruebas

### Qué hice:
- Se implementó la función ShowGoroutine.
- Se agregaron 10 goroutines ejecutándose de forma concurrente.
- Se utilizó la generación de tiempos aleatorios para simular diferentes retrasos.
- Se ejecutó el programa utilizando:

go run src/main.go

### Qué problema encontré:
- La salida de las goroutines no aparecía en orden numérico.

### Cómo lo resolví:
- Se verificó que el comportamiento era correcto debido a que las goroutines trabajan de forma concurrente.
- Se mantuvo una espera al final del programa para permitir que todas las tareas finalizaran.

### Próximo paso:
- Mejorar la documentación del proyecto.


---

## 20/07/2026 - Taller 20: Goroutines - Fase 3: Documentación

### Qué hice:
- Se creó el README.md propio del Taller 20.
- Se documentó:
  - Objetivo del proyecto.
  - Descripción del funcionamiento.
  - Instrucciones de ejecución.
  - Decisiones de diseño.
- Se actualizó el README principal del repositorio.

### Qué problema encontré:
- El archivo del README tenía un nombre diferente al solicitado por la estructura del proyecto.

### Cómo lo resolví:
- Se utilizó Git para cambiar el nombre del archivo:
  - Readme.md
  - README.md

### Próximo paso:
- Continuar mejorando el código y documentación.


---

## 20/07/2026 - Taller 20: Goroutines - Fase 4: Comentarios y organización del código

### Qué hice:
- Se agregaron comentarios explicativos dentro del archivo main.go.
- Se documentó el funcionamiento de la función ShowGoroutine.
- Se explicaron las partes principales del programa:
  - Generación de retrasos aleatorios.
  - Creación de goroutines.
  - Tiempo de espera del programa.

### Qué problema encontré:
- Era necesario mejorar la comprensión del código para facilitar su revisión.

### Cómo lo resolví:
- Se agregaron comentarios dentro del código explicando la función de cada bloque.

### Próximo paso:
- Revisar el estado final del proyecto y realizar pruebas adicionales.


---

## 20/07/2026 - Taller 20: Goroutines - Fase 5: Control de versiones

### Qué hice:
- Se realizaron commits progresivos utilizando mensajes descriptivos.
- Se subieron los avances al repositorio remoto de GitHub.
- Se verificó el historial de commits mediante Git.

### Qué problema encontré:
- Al inicio no se tenía una estructura organizada del repositorio.

### Cómo lo resolví:
- Se reorganizaron las carpetas siguiendo el instructivo del proyecto.
- Se mantuvo un historial de cambios mediante commits separados.

### Próximo paso:
- Finalizar la revisión del repositorio antes de la entrega.