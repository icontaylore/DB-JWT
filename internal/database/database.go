package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

type Database struct {
	db *sqlx.DB
}

func ConnectDb() *sqlx.DB {
	d := &Database{}
	db, err := sqlx.Connect("postgres", "user=postgres port=5432 password=1 dbname=postgres sslmode=disable")
	if err != nil {
		log.Printf("[ОШИБКА СТАРТА DB]", err)
	}

	fmt.Println("[DB START]")
	d.db = db
	return d.db
}
