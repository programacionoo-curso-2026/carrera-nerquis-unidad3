package dataaccess

import (
	"database/sql"
	"log"

	"github.com/glebarez/sqlite" //[cite: 2]
)

func InitDB() *sql.DB {
	// Usamos sqlite.Open tal como indica el material de clase
	db, err := sql.Open(sqlite.DriverName, "competenciasdocentes.db") //[cite: 2]
	if err != nil {
		log.Fatal(err) //[cite: 2]
	}

	err = db.Ping() //[cite: 2]
	if err != nil {
		log.Fatal(err) //[cite: 2]
	}

	log.Println("¡Conectado a SQLite con éxito!") //[cite: 2]
	return db
}
