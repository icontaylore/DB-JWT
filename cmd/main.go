package main

import (
	"github.com/gin-gonic/gin"
	"postgresandjwt/internal/database"
	"postgresandjwt/internal/handler"
	"postgresandjwt/internal/server"
	"postgresandjwt/internal/sescoock"
)

func main() {
	// INIT GIN
	router := gin.Default()
	router.LoadHTMLGlob("../template/*")
	// создаем сессии и куки
	sescoock.GoCoockieAndSession(router)
	// CREATE DB and CONNECT //
	db := database.ConnectDb()
	// HANDLER REG //
	router.GET("/register", handler.HandlerRegGet)
	router.POST("/register", handler.HandlerRegPost(db))
	// HANDLER LOG //
	router.GET("/login", handler.HandlerLogGet)
	router.POST("/login", handler.HandlerLogPost(db))

	s := server.CreateServer(router)
	s.Run()
}
