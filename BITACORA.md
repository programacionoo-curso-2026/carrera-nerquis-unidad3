# BITÁCORA DE AVANCES

## 18/07/2026 – Taller 20 – Fase 1: Organización inicial del proyecto

### Qué hice:
- Creé la carpeta correspondiente al Taller 20.
- Configuré la estructura inicial del programa en Go.
- Preparé el archivo principal (`main.go`) para implementar el uso de goroutines.

### Qué problema encontré:
- Al ejecutar el programa aparecieron errores relacionados con la organización del código y la estructura del proyecto.

### Cómo lo resolví:
- Reorganicé los archivos y corregí la estructura del proyecto para que el código compilara correctamente.

### Próximo paso:
- Implementar las goroutines y verificar su ejecución.

---

## 18/07/2026 – Taller 20 – Fase 2: Implementación de goroutines

### Qué hice:
- Implementé la lógica utilizando goroutines.
- Probé la ejecución concurrente del programa.
- Verifiqué que los resultados fueran los esperados.

### Qué problema encontré:
- La salida del programa no siempre aparecía en el orden esperado debido a la ejecución concurrente.

### Cómo lo resolví:
- Revisé la sincronización de la ejecución y validé el funcionamiento del programa mediante varias pruebas.

### Próximo paso:
- Iniciar la configuración del acceso a base de datos.

---

## 20/07/2026 – Taller 21 – Fase 1: Configuración de DataAccess

### Qué hice:
- Creé la carpeta del Taller 21.
- Configuré el módulo Go.
- Instalé el driver SQLite.
- Implementé el paquete `dataaccess`.

### Qué problema encontré:
- La conexión con SQLite no se establecía correctamente durante las primeras pruebas.

### Cómo lo resolví:
- Verifiqué la instalación del driver y corregí la inicialización de la base de datos utilizando `sql.Open()` y `Ping()`.

### Próximo paso:
- Integrar la conexión con el programa principal.

---

## 20/07/2026 – Taller 21 – Fase 2: Pruebas de conexión

### Qué hice:
- Integré el paquete `dataaccess` con `main.go`.
- Probé la apertura y cierre de la base de datos.
- Verifiqué que la conexión funcionara correctamente.

### Qué problema encontré:
- Existían errores en las rutas de importación del proyecto.

### Cómo lo resolví:
- Actualicé las importaciones para que coincidieran con el nombre definido en `go.mod`.

### Próximo paso:
- Continuar con el siguiente taller.

---

## 22/07/2026 – Taller 22 – Fase 1: Desarrollo del programa

### Qué hice:
- Creé la carpeta correspondiente al Taller 22.
- Implementé la lógica solicitada utilizando goroutines.
- Organicé el código para facilitar su mantenimiento.

### Qué problema encontré:
- Algunas funciones no se ejecutaban en el orden esperado debido a la concurrencia.

### Cómo lo resolví:
- Revisé la secuencia de ejecución y realicé pruebas hasta obtener el comportamiento esperado.

### Próximo paso:
- Documentar el funcionamiento del taller.

---

## 22/07/2026 – Taller 22 – Fase 2: Validación final

### Qué hice:
- Verifiqué el correcto funcionamiento del programa.
- Revisé el código y actualicé el README del taller.

### Qué problema encontré:
- Fue necesario realizar pequeños ajustes en la estructura del programa para mejorar su organización.

### Cómo lo resolví:
- Reorganicé el código y confirmé nuevamente la ejecución del programa.

### Próximo paso:
- Comenzar el desarrollo del Deber 4.

---

## 24/07/2026 – Deber 4 – Fase 1: Implementación del DocenteDAO

### Qué hice:
- Creé la carpeta `deber4-docente_dao`.
- Implementé la estructura `DocenteDAO`.
- Desarrollé el constructor `NewDocenteDAO()`.
- Implementé el método `CreateTable()` para crear la tabla de docentes.

### Qué problema encontré:
- Se presentaron errores durante las primeras pruebas de creación de la tabla.

### Cómo lo resolví:
- Revisé la sentencia SQL y corregí la definición de los campos de la tabla.

### Próximo paso:
- Implementar la inserción de docentes.

---

## 25/07/2026 – Deber 4 – Fase 2: Inserción de información

### Qué hice:
- Implementé el método `Insert()`.
- Probé la inserción de registros desde `main.go`.
- Verifiqué que los datos quedaran almacenados correctamente en SQLite.

### Qué problema encontré:
- Se generaban errores al intentar insertar registros repetidos.

### Cómo lo resolví:
- Revisé el uso de la clave primaria y del campo `email` para evitar duplicados.

### Próximo paso:
- Implementar las búsquedas requeridas para el siguiente deber.

---

## 28/07/2026 – Deber 5 – Fase 1: Búsqueda de docentes

### Qué hice:
- Creé la carpeta `deber5-docente_dao`.
- Implementé el método `GetByID()`.
- Implementé el método `GetByEmail()`.
- Reutilicé la estructura desarrollada en el Deber 4.

### Qué problema encontré:
- Era necesario controlar los casos en los que la consulta no encontraba registros.

### Cómo lo resolví:
- Implementé el manejo de `sql.ErrNoRows` para retornar mensajes de error adecuados cuando el docente no existe.

### Próximo paso:
- Integrar ambos métodos en `main.go`.

---

## 28/07/2026 – Deber 5 – Fase 2: Integración y pruebas

### Qué hice:
- Actualicé `main.go` para crear la tabla, insertar docentes y realizar búsquedas por ID y correo electrónico.
- Probé el funcionamiento completo del DAO.
- Preparé la documentación del README para registrar las evidencias de ejecución.

### Qué problema encontré:
- Fue necesario verificar que los datos de prueba coincidieran con las consultas realizadas.

### Cómo lo resolví:
- Ajusté los registros utilizados durante las pruebas y confirmé que las búsquedas devolvieran la información correcta.

### Próximo paso:
- Realizar los commits finales, actualizar la documentación y publicar la versión definitiva del proyecto en GitHub.