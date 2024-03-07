package services

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // PostgreSQL driver
)

func ConnectToDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("error connecting to db:%v", err)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
		return nil, err
	}

	log.Println("Connected to postgres successfully")

	return db, nil
}
