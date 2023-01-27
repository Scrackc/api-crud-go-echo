package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Scrackc/api-echo-crud/model"
	"github.com/labstack/echo/v4"
)

// person
type person struct {
	storage Storage
}

func newPerson(storage Storage) person {
	return person{storage}
}

func (p *person) create(c echo.Context) error {

	data := model.Person{}
	err := c.Bind(&data)
	if err != nil {
		response := newResponse(Error, "La persona no tiene una estructura correcta", nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	err = p.storage.Create(&data)
	if err != nil {
		response := newResponse(Error, "Ocurrio un error al crear la persona", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := newResponse(Message, "Persona creada", nil)
	return c.JSON(http.StatusCreated, response)
}

func (p *person) update(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := newResponse(Error, "El id debe ser un numúmero entero positivo", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	data := model.Person{}
	err = c.Bind(&data)

	if err != nil {
		response := newResponse(Error, "La persona no tiene una estructura correcta", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	err = p.storage.Update(id, &data)
	if err != nil {
		response := newResponse(Error, "Error al actualizar a la pesona", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}
	response := newResponse(Message, "Persona actualizada", nil)
	return c.JSON(http.StatusOK, response)
}

func (p *person) getAll(c echo.Context) error {

	resp, err := p.storage.GetAll()
	if err != nil {
		response := newResponse(Error, "Error al obtener las personas", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := newResponse(Message, "OK", resp)
	return c.JSON(http.StatusOK, response)
}

func (p *person) delete(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := newResponse(Error, "El id debe ser un numúmero entero positivo", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	err = p.storage.Delete(id)
	if errors.Is(err, model.ErrIdPersonDoesNotExists) {
		response := newResponse(Error, "Persona no existente", nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	if err != nil {
		response := newResponse(Error, "Error al eliminar la persona", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}
	response := newResponse(Message, "Persona eliminada", nil)
	return c.JSON(http.StatusOK, response)
}

func (p *person) getByID(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		respose := newResponse(Error, "El id debe ser de tipo entero positivo", nil)
		return c.JSON(http.StatusBadRequest, respose)
	}

	person, err := p.storage.GetByID(id)
	if errors.Is(err, model.ErrIdPersonDoesNotExists) {
		response := newResponse(Error, "Persona no existente", nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	if err != nil {
		response := newResponse(Error, "Error al obtener a la persona", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := newResponse(Message, "OK", person)
	return c.JSON(http.StatusOK, response)

}
