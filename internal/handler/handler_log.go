package handler

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func HandlerLogGet(c *gin.Context) {
	session := sessions.Default(c)
	errorMsg := session.Get("Error")
	if errorMsg != nil {
		c.HTML(200, "login.html", gin.H{
			"Error": errorMsg,
		})
	} else {
		c.HTML(200, "login.html", nil)
	}
}

func HandlerLogPost(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
