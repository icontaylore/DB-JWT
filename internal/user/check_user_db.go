package user

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

func CheckInDatabase(db *sqlx.DB, email string) bool {
	fmt.Println(email)
	var exists bool
	err := db.Get(&exists, "select exists(select 1 from users where email =$1)", email)
	if err != nil {
		log.Printf("[ОШИБКА СКАНИРОВАНИЯ]")
	}
	if !exists {
		return false
	} else {
		return true
	}
}
