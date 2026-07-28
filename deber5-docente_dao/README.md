# Deber 5 - Data Access Object (DAO) Docente

## Información del proyecto

**Asignatura:** Programación Orientada a Objetos  
**Unidad:** Unidad 3  
**Lenguaje:** Go  
**Base de datos:** SQLite  

**Estudiante:** Nerquis Carrera  
**Universidad:** Universidad Internacional del Ecuador  

---

# Objetivo

Implementar el patrón de diseño **Data Access Object (DAO)** utilizando el lenguaje Go y una base de datos SQLite.

El objetivo principal es separar la lógica de acceso a datos de la lógica principal del programa, permitiendo administrar la información de docentes mediante operaciones de persistencia.

El sistema permite:

- Crear la tabla de docentes automáticamente.
- Insertar registros de docentes.
- Buscar docentes mediante su identificador único.
- Buscar docentes mediante su correo electrónico.

---

# Descripción del proyecto

El proyecto implementa una estructura organizada mediante paquetes:

- **model:** contiene las entidades del sistema.
- **dataaccess:** administra la conexión con la base de datos SQLite.
- **dao:** contiene las operaciones de acceso y manipulación de datos.
- **main:** ejecuta las pruebas del funcionamiento del DAO.

Se utiliza SQLite como sistema gestor de base de datos debido a que permite trabajar con una base de datos local sin necesidad de configuraciones adicionales.

---

# Estructura del proyecto

```
deber5-docente_dao/

│
├── dao/
│   └── docente_dao.go
│
├── dataaccess/
│   └── dataaccess.go
│
├── model/
│   ├── docente.go
│   ├── competencia.go
│   └── evaluacion.go
│
├── main.go
├── go.mod
├── go.sum
└── README.md
```

---

# Descripción de paquetes

## Paquete model

Contiene las estructuras que representan las entidades utilizadas en el sistema.

### Docente

La entidad docente contiene los siguientes atributos:

- ID
- Nombre
- Email
- Departamento
- Cargo
- Años de antigüedad

---

## Paquete dataaccess

Este paquete permite establecer la conexión con SQLite mediante el paquete:

```
database/sql
```

y el driver:

```
github.com/glebarez/sqlite
```

La función principal es:

```go
InitDB()
```

Esta función:

- Abre la conexión con la base de datos.
- Verifica la disponibilidad de SQLite.
- Retorna la conexión para ser utilizada por el DAO.

---

## Paquete dao

Este paquete contiene la clase:

```
DocenteDAO
```

La cual implementa las operaciones de acceso a datos.

Métodos implementados:

### NewDocenteDAO()

Permite crear una instancia del DAO asociada a la conexión de base de datos.

---

### CreateTable()

Crea la tabla:

```
docentes
```

con los siguientes campos:

| Campo | Tipo |
|---|---|
| id | TEXT PRIMARY KEY |
| nombre | TEXT |
| email | TEXT UNIQUE |
| departamento | TEXT |
| cargo | TEXT |
| anios_antiguedad | INTEGER |

---

### Insert()

Permite almacenar nuevos docentes dentro de la base de datos.

Ejemplo:

```
D001 - Ana García
D002 - Carlos Ruiz
```

---

### GetByID()

Permite consultar un docente utilizando su identificador.

Ejemplo:

```
Buscar docente con ID D001
```

Resultado:

```
Ana García
```

---

### GetByEmail()

Permite consultar un docente utilizando su correo electrónico.

Ejemplo:

```
Buscar docente con email:
carlos.ruiz@email.com
```

Resultado:

```
Carlos Ruiz
```

---

# Requisitos para ejecutar

Antes de ejecutar el proyecto se necesita:

- Go instalado.
- Driver SQLite descargado.

Instalación del driver:

```bash
go get github.com/glebarez/sqlite
```

---

# Ejecución del programa

Ubicarse dentro de la carpeta:

```bash
cd deber5-docente_dao
```

Ejecutar:

```bash
go run main.go
```

---

# Resultado de ejecución

La ejecución correcta del programa genera la siguiente salida:

```
¡Conectado a SQLite con éxito!
Base de datos inicializada correctamente
Tabla docentes creada/verificada exitosamente
Docente D001 insertado exitosamente
Docente D002 insertado exitosamente
Docente encontrado:
&{ID:D001 Nombre:Ana García Email:ana.garcia@email.com Departamento:Informática Cargo:Profesora AniosAntiguedad:5}

Docente encontrado por email:
&{ID:D002 Nombre:Carlos Ruiz Email:carlos.ruiz@email.com Departamento:Matemáticas Cargo:Profesor AniosAntiguedad:3}
```

---

# Evidencias

Las evidencias del funcionamiento incluyen:

- Conexión exitosa con SQLite.
- Creación automática de la tabla docentes.
- Inserción de registros.
- Consulta mediante ID.
- Consulta mediante correo electrónico.

Se recomienda agregar capturas de pantalla de la ejecución dentro de la carpeta:

```
docs/
```

del repositorio principal.

---

# Decisiones de diseño

Para el desarrollo del proyecto se aplicaron los siguientes criterios:

- Separación de responsabilidades mediante paquetes.
- Uso del patrón DAO para manejar el acceso a datos.
- Uso de estructuras (`struct`) para representar entidades.
- Uso de SQLite para almacenamiento persistente.
- Manejo de errores mediante mensajes descriptivos.

---

# Control de versiones

El desarrollo del proyecto se realizó utilizando Git y GitHub.

Los cambios fueron registrados mediante commits progresivos siguiendo buenas prácticas:

- Commits con mensajes descriptivos.
- Registro de avances por etapas.
- Documentación del proceso de desarrollo.

---

# Conclusión

La implementación del patrón DAO permitió conectar una aplicación desarrollada en Go con una base de datos SQLite, logrando realizar operaciones de persistencia sobre la entidad Docente.

El proyecto demuestra la separación entre la lógica de negocio y el acceso a datos, facilitando el mantenimiento y escalabilidad del sistema.