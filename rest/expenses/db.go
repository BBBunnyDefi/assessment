package expenses

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB(url string) *sql.DB {
	var err error
	db, err = sql.Open("postgres", url)
	if err != nil {
		log.Fatal("Connect to database error:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Check database server:", err)
	}

	// create database
	sql := `
		CREATE TABLE IF NOT EXISTS expenses (
			id SERIAL PRIMARY KEY,
			title TEXT,
			amount FLOAT,
			note TEXT,
			tags TEXT[]
		);
	`
	_, err = db.Exec(sql)
	if err != nil {
		log.Fatal("Can't create table:", err)
	}

	log.Println("Database connection OK!!")

	return db
}
