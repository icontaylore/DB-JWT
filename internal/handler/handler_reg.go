package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"log"
	"postgresandjwt/internal/user"
)

func HandlerRegGet(c *gin.Context) {
	c.HTML(200, "register.html", nil)
}

func HandlerRegPost(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		u := &user.User{}
		if err := c.ShouldBind(&u); err != nil {
			log.Printf("[SHOULD BIND ERR]", err)
		}

		user.CheckInDatabase(db)
	}
}
