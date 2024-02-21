package http

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/marcin-kupiec/srv-numbers/numbers"
	"github.com/marcin-kupiec/srv-numbers/numbers/adapter/http/models"
)

//go:generate moq -out numbersServiceMock_test.go . numbersService
type numbersService interface {
	Get(ctx context.Context, number int64) (int64, int64, error)
}

func NewGetHandler(service numbersService) echo.HandlerFunc {
	return func(c echo.Context) error {
		number, err := strconv.ParseInt(c.Param("number"), 10, 64)
		if err != nil {
			c.Logger().Debugf("GetHandler - received invalid number parameter %d: %v", number, err)
			return errInvalidRequest
		}

		index, value, err := service.Get(c.Request().Context(), number)
		if err != nil {
			if errors.Is(err, numbers.ErrNumberNotFound) {
				return errNumberNotFound
			}
			c.Logger().Errorf("GetHandler - failed to get number index: %v", err)
			return errInternalServerError
		}

		return c.JSON(http.StatusOK, models.GetNumberResponse{
			ID:    &index,
			Value: &value,
		})
	}
}
