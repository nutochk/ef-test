package server

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nutochk/ef-test/internal/dto"
	"github.com/nutochk/ef-test/internal/models"
	"github.com/nutochk/ef-test/internal/repository"
)

// CreatePerson godoc
// @Summary create new record about person
// @Description Creates a new record with data enrichment from external APIs
// @Tags people
// @Accept  json
// @Produce  json
// @Param person body models.Person true "Personal data"
// @Success 200 {object} dto.PersonInfo
// @Failure 400 {string} string "Incorrect data format"
// @Failure 500 {string} string "Server error"
// @Router /api/people [post]
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

// UpdatePerson godoc
// @Summary update record about person
// @Description Updates the record of an existing person
// @Tags people
// @Accept  json
// @Produce  json
// @Param        id   path      int  true  "Person ID"
// @Param person body models.Person true "Personal data"
// @Success 200 {object} dto.PersonInfo
// @Failure 400 {string} string "Incorrect data format"
// @Failure 404 {string} string "Not found"
// @Failure 500 {string} string "Server error"
// @Router /api/people/{id} [put]
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

// DeletePerson godoc
// @Summary delete record about person
// @Description Delete the record of an existing person
// @Tags people
// @Accept  json
// @Produce  json
// @Param        id   path      int  true  "Person ID"
// @Success 204
// @Failure 400 {string} string "Incorrect data format"
// @Failure 404 {string} string "Not found"
// @Failure 500 {string} string "Server error"
// @Router /api/people/{id} [delete]
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

// GetPerson godoc
// @Summary get by id record about person
// @Description get the record of an existing person
// @Tags people
// @Accept  json
// @Produce  json
// @Param        id   path      int  true  "Person ID"
// @Success 200 {object} []models.PersonInfo
// @Failure 400 {string} string "Incorrect data format"
// @Failure 404 {string} string "Not found"
// @Failure 500 {string} string "Server error"
// @Router /api/people/{id} [get]
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

// GetPeople godoc
// @Summary Get a list of people with filtering
// @Description Returns a list of people with the ability to filter and paginate
// @Tags people
// @Accept  json
// @Produce  json
// @Param name query string false "Filter by name"
// @Param surname query string false "Filter by last name"
// @Param age_min query int false "Minimum age"
// @Param age_max query int false "Maximum age"
// @Param gender query string false "Gender filter (male/female)"
// @Param page query int false "Page number" default(1)
// @Param per_page query int false "Number of entries per page" default(10)
// @Success 200 {object} dto.PaginatedResponse
// @Failure 400 {object} map[string]string "Incorrect filtering parameters"
// @Failure 500 {string} string "Server error"
// @Router /api/people [get]
func (server *Server) getPeople(c *gin.Context) {
	var filters dto.PersonFilter
	if err := c.ShouldBindQuery(&filters); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid filters"})
		return
	}
	var pagination dto.Pagination
	if err := c.ShouldBindQuery(&pagination); err != nil {
		pagination = dto.Pagination{Page: 1, PerPage: 10}
	}
	if pagination.Page <= 0 {
		pagination.Page = 1
	}
	if pagination.PerPage <= 0 {
		pagination.PerPage = 10
	}

	response, err := server.service.GetPeople(&filters, &pagination)
	if err != nil {
		c.String(http.StatusInternalServerError, "failed to find people with filters")
		return
	}
	c.JSON(http.StatusOK, response)
}
