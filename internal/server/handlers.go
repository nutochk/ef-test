package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (server *Server) create(c *gin.Context) {
	fmt.Println("create method in server")
}
