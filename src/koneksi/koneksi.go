package koneksi

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Konek() (*sql.DB, error) {
	const userDB string = "fizi"
	const passDB string = "fizi12345"
	const hostDB string = "multimatics-mysql"
	const portDB string = "3306"
	const nameDB string = "multimatics"

	db, err := sql.Open("mysql", userDB+":"+passDB+"@tcp("+hostDB+":"+portDB+")/"+nameDB)
	if err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Database conection error: %v", err)
	}
	return db, nil
}
