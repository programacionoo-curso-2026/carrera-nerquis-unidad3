# Deber 4 - DataAccessObject para Docente (docente_dao)

## Datos del estudiante

**Estudiante:** Nerquis Carrera  
**Carrera:** Ingeniería en Sistemas de la Información  
**Universidad:** Universidad Internacional del Ecuador  
**Unidad:** Unidad 3  

---

# Descripción del proyecto

Este proyecto corresponde al desarrollo del patrón **DataAccessObject (DAO)** aplicado a la entidad **Docente**, utilizando el lenguaje de programación Go y una base de datos SQLite.

La finalidad del proyecto es implementar una estructura organizada para separar la conexión con la base de datos, el modelo de datos y las operaciones de acceso mediante un DAO.

---

# Objetivo

Desarrollar un DataAccessObject para la gestión de docentes utilizando Go y SQLite, permitiendo realizar operaciones de persistencia de datos mediante el paquete `database/sql`.

Las funcionalidades implementadas son:

- Conexión con una base de datos SQLite.
- Creación automática de la tabla docentes.
- Inserción de registros de docentes.
- Consulta de docentes mediante su ID.
- Consulta de docentes mediante su correo electrónico.

---

# Tecnologías utilizadas

- Lenguaje de programación: Go
- Base de datos: SQLite
- Librería utilizada:


github.com/glebarez/sqlite


- Paquete utilizado para conexión:


database/sql


---

# Estructura del proyecto


deber4-docente_dao
│
├── dao
│ └── docente_dao.go
│
├── dataaccess
│ └── dataaccess.go
│
├── model
│ └── docente.go
│
├── main.go
├── go.mod
└── go.sum


---

# Descripción de los componentes

## DataAccess

Archivo:


dataaccess/dataaccess.go


Este paquete permite realizar la conexión con la base de datos SQLite mediante la función `InitDB()`.

La conexión utiliza la base de datos:


competenciasdocentes.db


y verifica que la conexión se encuentre disponible antes de retornar la instancia de la base de datos.

---

## Model

Archivo:


model/docente.go


Contiene la estructura que representa la entidad Docente.

Los atributos definidos son:

- ID
- Nombre
- Email
- Departamento
- Cargo
- AniosAntiguedad

Esta estructura permite manejar la información de los docentes dentro del programa.

---

## DAO

Archivo:


dao/docente_dao.go


Este paquete implementa el patrón **DataAccessObject** para manejar las operaciones relacionadas con la tabla docentes.

Funciones implementadas:

### NewDocenteDAO()

Permite crear una instancia del DAO utilizando la conexión con SQLite.

### CreateTable()

Crea la tabla docentes si esta no existe.

Estructura creada:

```sql
CREATE TABLE IF NOT EXISTS docentes (
    id TEXT PRIMARY KEY,
    nombre TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    departamento TEXT,
    cargo TEXT,
    anios_antiguedad INTEGER DEFAULT 0
);
Insert()

Permite insertar nuevos docentes en la base de datos.

GetByID()

Permite obtener un docente utilizando su identificador único.

GetByEmail()

Permite obtener un docente mediante su correo electrónico.

Ejecución del programa

Para ejecutar el proyecto se debe ingresar a la carpeta:

deber4-docente_dao

y ejecutar el siguiente comando:

go run .
Evidencia de ejecución

La ejecución del programa fue realizada correctamente obteniendo los siguientes resultados:

2026/07/28 10:58:44 ¡Conectado a SQLite con éxito!
2026/07/28 10:58:44 Base de datos inicializada correctamente
2026/07/28 10:58:44 Tabla docentes creada/verificada exitosamente
2026/07/28 10:58:44 Docente D001 insertado exitosamente
2026/07/28 10:58:44 Docente D002 insertado exitosamente
2026/07/28 10:58:44 Docente encontrado: &{ID:D001 Nombre:Ana Garcia Email:ana.garcia@email.com Departamento:Informatica Cargo:Profesora AniosAntiguedad:5}
2026/07/28 10:58:44 Docente encontrado por email: &{ID:D002 Nombre:Carlos Ruiz Email:carlos.ruiz@email.com Departamento:Matematicas Cargo:Profesor AniosAntiguedad:3}
Validación de funcionamiento

Mediante la ejecución del programa se comprobó que:

La conexión con SQLite fue exitosa.
La tabla docentes fue creada correctamente.
Los registros de docentes fueron insertados.
La consulta por ID funcionó correctamente.
La consulta por correo electrónico funcionó correctamente.
Conclusión

El desarrollo del deber permitió implementar correctamente el patrón DataAccessObject (DAO) en Go, separando la conexión a la base de datos, el modelo de información y la lógica de acceso a datos.

El programa cumple con los requerimientos solicitados, utilizando SQLite como sistema de almacenamiento y permitiendo gestionar información de docentes mediante operaciones de persistencia y consulta.


Este ya está listo para guardarlo como:

```text
deber4-docente_dao/README.md