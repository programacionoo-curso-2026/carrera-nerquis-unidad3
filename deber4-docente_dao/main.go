package main

import (
	"log"

	"github.com/programacionoo-curso-2026/carrera-nerquis-unidad3/deber4-docente_dao/dao"
	"github.com/programacionoo-curso-2026/carrera-nerquis-unidad3/deber4-docente_dao/dataaccess"
	"github.com/programacionoo-curso-2026/carrera-nerquis-unidad3/deber4-docente_dao/model"
)

func main() {

	// Inicializar la base de datos
	db := dataaccess.InitDB()
	defer db.Close()

	log.Println("Base de datos inicializada correctamente")

	// Crear el DAO
	docenteDAO := dao.NewDocenteDAO(db)

	// Crear la tabla
	if err := docenteDAO.CreateTable(); err != nil {
		log.Fatalf("Error al crear tabla: %v", err)
	}

	// Insertar docentes

	docente1 := &model.Docente{
		ID:              "D001",
		Nombre:          "Ana Garcia",
		Email:           "ana.garcia@email.com",
		Departamento:    "Informatica",
		Cargo:           "Profesora",
		AniosAntiguedad: 5,
	}

	err := docenteDAO.Insert(docente1)

	if err != nil {
		log.Printf("Error al insertar docente: %v", err)
	}

	docente2 := &model.Docente{
		ID:              "D002",
		Nombre:          "Carlos Ruiz",
		Email:           "carlos.ruiz@email.com",
		Departamento:    "Matematicas",
		Cargo:           "Profesor",
		AniosAntiguedad: 3,
	}

	err = docenteDAO.Insert(docente2)

	if err != nil {
		log.Printf("Error al insertar docente: %v", err)
	}

	// Buscar por ID

	docente, err := docenteDAO.GetByID("D001")

	if err != nil {
		log.Printf("Error al buscar docente: %v", err)

	} else {

		log.Printf("Docente encontrado: %+v", docente)
	}

	// Buscar por Email

	docenteEmail, err := docenteDAO.GetByEmail("carlos.ruiz@email.com")

	if err != nil {

		log.Printf("Error al buscar por email: %v", err)

	} else {

		log.Printf("Docente encontrado por email: %+v", docenteEmail)
	}
}
