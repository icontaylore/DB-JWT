package sescoock

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func GoCoockieAndSession(router *gin.Engine) {
	store := cookie.NewStore([]byte("next_time_change"))
	router.Use(sessions.Sessions("mysession", store))
}
