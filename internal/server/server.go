package server

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nutochk/ef-test/internal/service"
)

type Server struct {
	engine     *gin.Engine
	service    service.Service
	httpServer *http.Server
}

func New(service service.Service) *Server {
	e := gin.Default()
	s := &Server{
		engine:  e,
		service: service,
		httpServer: &http.Server{
			Handler: e,
		},
	}
	s.registerRouters()
	return s
}

func (s *Server) registerRouters() {
	api := s.engine.Group("/api")
	{
		api.POST("/people", s.create)
		api.PUT("/people/:id", s.update)
		api.DELETE("/people/:id", s.delete)
		api.GET("people/:id", s.getById)
		api.GET("/people", s.getPeople)
	}
}

func (s *Server) Run(port int) error {
	s.httpServer.Addr = ":" + strconv.Itoa(port)
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
