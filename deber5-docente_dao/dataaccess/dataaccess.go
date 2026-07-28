package dataaccess

import (
	"database/sql"
	"log"

	_ "github.com/glebarez/sqlite"
)

// InitDB establece la conexión con la base de datos SQLite.
func InitDB() *sql.DB {

	db, err := sql.Open("sqlite", "competenciasdocentes.db")
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
