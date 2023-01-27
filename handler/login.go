package handler

import (
	"net/http"

	"github.com/Scrackc/api-echo-crud/auth"
	"github.com/Scrackc/api-echo-crud/model"
	"github.com/labstack/echo/v4"
)

type login struct {
	storage Storage
}

func newLogin(s Storage) login {
	return login{s}
}

func (l *login) login(c echo.Context) error {

	data := model.Login{}
	err := c.Bind(&data)
	if err != nil {
		response := newResponse(Error, "estructura no valida", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	if !isLoginValid(&data) {
		response := newResponse(Error, "credenciales no validas", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	token, err := auth.GenerateToken(&data)
	if err != nil {
		response := newResponse(Error, "No se pudo generar el token", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}
	dataToken := map[string]string{"token": token}
	response := newResponse(Message, "OK", dataToken)
	return c.JSON(http.StatusBadRequest, response)

}

func isLoginValid(data *model.Login) bool {
	return data.Email == "ed@mail.com" && data.Password == "123456"
}
