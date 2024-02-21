package http

import (
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"numbers/numbers/adapter/http/models"
)

type numbersService interface {
	Get(ctx context.Context, number int) (int64, int64, error)
}

func NewGetter(service numbersService) echo.HandlerFunc {
	return func(c echo.Context) error {
		number, err := strconv.Atoi(c.Param("number"))
		if err != nil {
			return echo.ErrBadRequest
		}

		index, value, err := service.Get(c.Request().Context(), number)
		if err != nil {
			return echo.ErrInternalServerError
		}

		return c.JSON(http.StatusOK, models.GetNumberResponse{
			ID:           &index,
			Value:        &value,
			ErrorMessage: "",
		})
	}
}
