package user

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

func CheckInDatabase(db *sqlx.DB) {
	rows, err := db.Query("select * from users")
	if err != nil {
		log.Printf("[нельзя достать юзеров]", err)
	}

	for rows.Next() {
		u := User{}
		rows.Scan(&u.Id, &u.Email, &u.Password)
		fmt.Println(u)
	}
}
