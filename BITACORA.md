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
- Se implementó la primera versión del programa utilizando goroutines en Go.

### Qué problema encontré:
- El repositorio inicialmente fue clonado dentro de una carpeta incorrecta, generando una estructura duplicada.
- El código proporcionado originalmente presentaba errores de sintaxis que impedían su ejecución.

### Cómo lo resolví:
- Se eliminó la copia incorrecta del repositorio y se realizó nuevamente la clonación desde la ubicación correcta.
- Se corrigieron las instrucciones de importación, la impresión en pantalla y el manejo del tiempo para que el programa compilara correctamente.

### Próximo paso:
- Ejecutar el programa y verificar el funcionamiento de las goroutines.

---

## 20/07/2026 - Taller 20: Goroutines - Fase 2: Implementación y pruebas

### Qué hice:
- Se implementó la función ShowGoroutine.
- Se agregaron 10 goroutines ejecutándose de forma concurrente.
- Se utilizó la generación de tiempos aleatorios para simular diferentes retrasos.
- Se realizaron pruebas ejecutando el programa con:

go run src/main.go

### Qué problema encontré:
- La salida de las goroutines no aparecía en orden numérico.

### Cómo lo resolví:
- Se verificó que el comportamiento era correcto debido a que las goroutines se ejecutan de manera concurrente y cada una finaliza en un momento diferente.

### Próximo paso:
- Documentar el funcionamiento del proyecto.

---

## 20/07/2026 - Taller 20: Goroutines - Fase 3: Documentación

### Qué hice:
- Se creó el README.md del Taller 20.
- Se documentó el objetivo del proyecto.
- Se explicó el funcionamiento del programa.
- Se agregaron instrucciones para ejecutar el código.
- Se documentaron las decisiones de diseño.
- Se actualizó el README principal del repositorio.

### Qué problema encontré:
- El archivo README tenía un nombre diferente al solicitado por la estructura del proyecto.

### Cómo lo resolví:
- Se renombró correctamente el archivo utilizando Git para mantener el historial de cambios.

### Próximo paso:
- Mejorar la organización y legibilidad del código.

---

## 20/07/2026 - Taller 20: Goroutines - Fase 4: Comentarios y organización del código

### Qué hice:
- Se agregaron comentarios explicativos dentro del archivo main.go.
- Se documentó el funcionamiento de la función ShowGoroutine.
- Se explicó la generación de retrasos aleatorios.
- Se documentó la creación de las goroutines.
- Se explicó el funcionamiento general del programa.

### Qué problema encontré:
- Era necesario mejorar la comprensión del código para facilitar su revisión.

### Cómo lo resolví:
- Se añadieron comentarios descriptivos en cada parte importante del programa.

### Próximo paso:
- Revisar el historial de cambios y preparar la entrega.

---

## 20/07/2026 - Taller 20: Goroutines - Fase 5: Control de versiones

### Qué hice:
- Se realizaron múltiples commits utilizando mensajes descriptivos.
- Se subieron los cambios al repositorio remoto de GitHub.
- Se verificó el historial de commits.
- Se organizó correctamente la estructura del repositorio.

### Qué problema encontré:
- Al inicio no se tenía una estructura organizada del repositorio y algunos archivos requerían ajustes.

### Cómo lo resolví:
- Se reorganizaron las carpetas siguiendo el instructivo del proyecto.
- Se corrigieron los nombres de los archivos y se mantuvo un historial de cambios mediante commits separados.

### Próximo paso:
- Optimizar la sincronización de las goroutines.

---

## 20/07/2026 - Taller 20: Goroutines - Fase 6: Sincronización con WaitGroup

### Qué hice:
- Se reemplazó el uso de time.Sleep en la función principal por sync.WaitGroup.
- Se implementó un contador para esperar la finalización de las 10 goroutines.
- Se agregó un mensaje indicando que todas las goroutines finalizaron correctamente.
- Se realizaron nuevas pruebas para comprobar el correcto funcionamiento del programa.

### Qué problema encontré:
- El uso de time.Sleep no garantizaba que todas las goroutines terminaran antes de finalizar el programa, ya que dependía de un tiempo fijo de espera.

### Cómo lo resolví:
- Se implementó sync.WaitGroup para sincronizar la ejecución de todas las goroutines.
- Cada goroutine notifica cuando finaliza su ejecución y el programa espera hasta que todas concluyan antes de terminar.

### Próximo paso:
- Revisar la documentación, verificar la estructura del repositorio y confirmar que el proyecto esté listo para su entrega.

```markdown
### [21/07/2026] — Taller 21, Fase 1: Entidades y Conexión SQLite
- Qué hice: Creación del paquete model con la entidad Docente, configuración del paquete dataaccess con SQLite y prueba exitosa en el main.go.
- Qué problema encontré: Error de importación no utilizada con el driver de SQLite.
- Cómo lo resolví: Se actualizó la función `sql.Open` utilizando explícitamente `sqlite.DriverName` como se indicó en los materiales de clase.
- Próximo paso: Realizar los commits incrementales correspondientes.