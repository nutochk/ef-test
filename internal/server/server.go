package server

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nutochk/ef-test/internal/service"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	engine     *gin.Engine
	service    service.Service
	httpServer *http.Server
}

// @title People API
// @version 1.0
// @description API for work with information about people
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8084
// @BasePath /api
// @schemes http

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
	s.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
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
