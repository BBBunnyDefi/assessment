package expenses

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB(url string) *sql.DB {
	// url := "postgres://root:root@localhost:5432/v1assessment?sslmode=disable"
	var err error
	db, err = sql.Open("postgres", url)
	if err != nil {
		log.Fatal("Connect to database error:", err)
	}

	log.Println("Database connection OK!!")

	return db
}
