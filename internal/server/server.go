package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func CreateServer(router *gin.Engine) *Server {
	return &Server{&http.Server{
		Addr:         "localhost:8080",
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
		Handler:      router,
	}}
}

func (s *Server) Run() {
	fmt.Println("[START SERVER]")
	if err := s.httpServer.ListenAndServe(); err != nil {
		log.Printf("[SERVER DON LISTENEN]")
	}
}
