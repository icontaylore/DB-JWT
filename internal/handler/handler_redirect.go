package handler

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SomeHandler(c *gin.Context, err string) {
	// получаем сессию
	session := sessions.Default(c)
	// ставим и сохраняем ошибку
	session.Set("Error", err)
	session.Save()
	c.Redirect(http.StatusSeeOther, "/login")
}
