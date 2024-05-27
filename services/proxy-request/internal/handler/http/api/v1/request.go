package httpv1

import (
	"errors"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"github.com/sazonovItas/proxy-manager/proxy-request/internal/adapter"
)

type requestById struct {
	Id string `param:"id" validate:"required"`
}

func (h *Handler) RequestById(c echo.Context) error {
	var params requestById
	if err := c.Bind(&params); err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	if err := c.Validate(&params); err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	id, err := uuid.Parse(params.Id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "bad format of id")
	}

	request, err := h.requestUsc.Request(c.Request().Context(), id)
	if err != nil {
		switch {
		case errors.Is(err, adapter.ErrRequestNotFound):
			return c.JSON(http.StatusNotFound, "no requests found")
		default:
			return c.JSON(http.StatusInternalServerError, "internal error")
		}
	}

	return c.JSON(http.StatusOK, request)
}

type requestsByTimestamp struct {
	From int64 `query:"timestamp" validate:"required"`
	To   int64 `query:"to"        validate:"required"`
}

func (h *Handler) RequestsByTimestamp(c echo.Context) error {
	var params requestsByTimestamp
	if err := c.Bind(&params); err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	if err := c.Validate(&params); err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	requests, err := h.requestUsc.Timestamp(
		c.Request().Context(),
		time.Unix(params.From, 0),
		time.Unix(params.To, 0),
	)
	if err != nil {
		switch {
		case errors.Is(err, adapter.ErrRequestNotFound):
			return c.JSON(http.StatusNotFound, "no requests found")
		default:
			return c.JSON(http.StatusInternalServerError, "internal error")
		}
	}

	return c.JSON(http.StatusOK, requests)
}

type requestsById struct {
	Id string `param:"id" validate:"required"`
	requestsByTimestamp
}

func (h *Handler) initRequestRoutes(api *echo.Group) {
	request := api.Group("/request")
	{
		request.GET("/", h.RequestsByTimestamp)
		request.GET("/:id", h.RequestById)
	}
}
