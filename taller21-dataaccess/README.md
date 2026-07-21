# Taller 21 - Entidades del proyecto y DAO para registro de información

## Objetivo
Implementar las entidades del proyecto en el paquete `model`, configurar la persistencia de datos relacionales con SQLite en el paquete `dataaccess`, y establecer la estructura base para el registro y manejo de información según los lineamientos del curso.

## Estructura del Proyecto
El taller se encuentra organizado de forma modular de la siguiente manera:
- `model/`: Contiene las estructuras de datos (entidades o clases) con sus respectivos atributos.
- `dataaccess/`: Contiene la configuración y la conexión a la base de datos utilizando SQLite.
- `main.go`: Archivo principal que ejecuta la inicialización de la base de datos y la gestión de entidades[cite: 2].

## Instrucciones de Ejecución
1. Asegúrate de estar ubicado en la carpeta del taller (`taller21-dataaccess`).
2. Ejecuta el programa utilizando el siguiente comando en tu terminal:
   ```bash
   go run main.go