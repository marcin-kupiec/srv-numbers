package http

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/marcin-kupiec/srv-numbers/numbers/adapter/http/models"
)

var (
	errNumberNotFound      = fmt.Errorf("number not found")
	errInvalidRequest      = fmt.Errorf("invalid request")
	errInternalServerError = fmt.Errorf("internal server error")
)

func HandleError(err error, c echo.Context) {
	if echoError, ok := err.(*echo.HTTPError); ok {
		_ = c.JSON(echoError.Code, echoError)
		return
	}
	if errors.Is(err, errInvalidRequest) {
		_ = c.JSON(http.StatusBadRequest, &models.Error{Message: err.Error()})
		return
	}
	if errors.Is(err, errNumberNotFound) {
		_ = c.JSON(http.StatusNotFound, &models.Error{Message: err.Error()})
		return
	}
	if errors.Is(err, errInternalServerError) {
		_ = c.JSON(http.StatusInternalServerError, &models.Error{Message: err.Error()})
		return
	}

	_ = c.JSON(http.StatusInternalServerError, echo.ErrInternalServerError)
}
