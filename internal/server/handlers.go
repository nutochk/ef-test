package server

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nutochk/ef-test/internal/models"
	"github.com/nutochk/ef-test/internal/repository"
)

func (server *Server) create(c *gin.Context) {
	var person models.Person
	if err := c.ShouldBindJSON(&person); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	p, err := server.service.Create(&person)
	if err != nil {
		c.String(http.StatusInternalServerError, "failed to create person")
		return
	}
	c.JSON(http.StatusOK, p)
}

func (server *Server) update(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	var person models.Person
	if err := c.ShouldBindJSON(&person); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	pi, err := server.service.Update(id, &person)
	if err != nil {
		if errors.Is(err, repository.ErrNotExist) {
			c.String(http.StatusNotFound, err.Error())
			return
		}
		c.String(http.StatusInternalServerError, "failed to update person")
		return
	}
	c.JSON(http.StatusOK, pi)
}

func (server *Server) delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	err = server.service.Delete(id)
	if err != nil {
		if errors.Is(err, repository.ErrNotExist) {
			c.String(http.StatusNotFound, err.Error())
			return
		}
		c.String(http.StatusInternalServerError, "failed to delete person")
		return
	}
	c.Writer.WriteHeader(http.StatusNoContent)
}

func (server *Server) getById(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	pi, err := server.service.GetById(id)
	if err != nil {
		if errors.Is(err, repository.ErrNotExist) {
			c.String(http.StatusNotFound, err.Error())
			return
		}
		c.String(http.StatusInternalServerError, "failed to get person")
		return
	}
	c.JSON(http.StatusOK, pi)
}
