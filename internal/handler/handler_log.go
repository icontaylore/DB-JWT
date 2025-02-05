package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func HandlerLogGet(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

func HandlerLogPost(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
