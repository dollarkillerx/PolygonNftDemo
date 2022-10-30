package server

import (
	"github.com/dollarkillerx/PolygonNftDemo/internal/conf"
	"github.com/dollarkillerx/PolygonNftDemo/internal/middleware"
	"github.com/gin-gonic/gin"

	"sync"
)

type Server struct {
	app *gin.Engine

	mu    sync.Mutex
	cache map[string]string // token: address
}

func NewServer() *Server {
	ser := &Server{
		app:   gin.New(),
		cache: map[string]string{},
	}

	go ser.alchemy()

	return ser
}

func (s *Server) Run() error {
	s.app.Use(middleware.SetBasicInformation())
	s.app.Use(middleware.Cors())
	s.app.Use(middleware.HttpRecover())

	s.router()

	return s.app.Run(conf.CONF.ListenAddr)
}
