package http

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/assert"

	"github.com/marcin-kupiec/srv-numbers/numbers"
)

func TestGetHandler(t *testing.T) {
	e := echo.New()
	e.HTTPErrorHandler = HandleError
	e.Logger.SetLevel(log.OFF)

	t.Run("should handle get request and return 200", func(t *testing.T) {
		numbersSvc := &numbersServiceMock{
			GetFunc: func(ctx context.Context, number int64) (int64, int64, error) {
				return 4, 123, nil
			},
		}

		handler := NewGetHandler(numbersSvc)
		SetRoutes(e, handler)

		req := httptest.NewRequest(http.MethodGet, "/endpoint/123", nil)

		rr := httptest.NewRecorder()
		e.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, `{"id":4,"value":123}`, strings.Trim(rr.Body.String(), "\n"))
	})

	t.Run("should get 400 code for bad number param", func(t *testing.T) {
		handler := NewGetHandler(nil)
		SetRoutes(e, handler)

		req := httptest.NewRequest(http.MethodGet, "/endpoint/abc", nil)

		rr := httptest.NewRecorder()
		e.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, `{"message":"invalid request"}`, strings.Trim(rr.Body.String(), "\n"))
	})

	t.Run("should get 404 code for number not found", func(t *testing.T) {
		numbersSvc := &numbersServiceMock{
			GetFunc: func(ctx context.Context, number int64) (int64, int64, error) {
				return 0, 0, numbers.ErrNumberNotFound
			},
		}

		handler := NewGetHandler(numbersSvc)
		SetRoutes(e, handler)

		req := httptest.NewRequest(http.MethodGet, "/endpoint/123", nil)

		rr := httptest.NewRecorder()
		e.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code)
		assert.Equal(t, `{"message":"number not found"}`, strings.Trim(rr.Body.String(), "\n"))
	})

	t.Run("should get 500 code for internal server error", func(t *testing.T) {
		numbersSvc := &numbersServiceMock{
			GetFunc: func(ctx context.Context, number int64) (int64, int64, error) {
				return 0, 0, assert.AnError
			},
		}

		handler := NewGetHandler(numbersSvc)
		SetRoutes(e, handler)

		req := httptest.NewRequest(http.MethodGet, "/endpoint/123", nil)

		rr := httptest.NewRecorder()
		e.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Equal(t, `{"message":"internal server error"}`, strings.Trim(rr.Body.String(), "\n"))
	})
}
