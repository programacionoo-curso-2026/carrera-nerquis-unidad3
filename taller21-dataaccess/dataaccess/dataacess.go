package dataaccess

import (
	"database/sql"
	"log"

	"github.com/glebarez/sqlite"
)

func InitDB() *sql.DB {
	// Usamos sqlite.Open tal como indica el material de clase
	db, err := sql.Open(sqlite.DriverName, "competenciasdocentes.db")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("¡Conectado a SQLite con éxito!")
	return db
}
