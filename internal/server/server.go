package server

import (
	"github.com/dollarkillerx/PolygonNftDemo/internal/conf"
	"github.com/dollarkillerx/PolygonNftDemo/internal/middleware"
	"github.com/gin-gonic/gin"
)

type Server struct {
	app *gin.Engine
}

func NewServer() *Server {
	ser := &Server{
		app: gin.New(),
	}

	return ser
}

func (s *Server) Run() error {
	s.app.Use(middleware.SetBasicInformation())
	s.app.Use(middleware.Cors())
	s.app.Use(middleware.HttpRecover())

	s.router()

	return s.app.Run(conf.CONF.ListenAddr)
}
