package main

import (
	"competenciasdocentes/dataaccess"
	"competenciasdocentes/model"
	"log"
)

func main() {
	// 1. Inicializar la base de datos
	db := dataaccess.InitDB()
	defer db.Close() // Importante: cerrar la conexión al final

	log.Println("Base de datos inicializada correctamente")

	// 2. Crear una instancia de prueba para la entidad Docente
	docente1 := &model.Docente{
		ID:              "D001",
		Nombre:          "Ana García",
		Email:           "ana.garcia@email.com",
		Departamento:    "Informática",
		Cargo:           "Profesora",
		AniosAntiguedad: 5,
	}

	log.Printf("Docente creado en memoria: %+v", docente1)
}
