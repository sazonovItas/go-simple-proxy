package httpv1

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sazonovItas/proxy-manager/services/proxy-request/internal/entity"
)

type RequestUsecase interface {
	Save(ctx context.Context, r *entity.Request) error
	Request(ctx context.Context, id uuid.UUID) (*entity.Request, error)
	Timestamp(ctx context.Context, from time.Time, to time.Time) ([]entity.Request, error)
}

type Handler struct {
	requestUsc RequestUsecase
}

func NewHandler(requestUsc RequestUsecase) *Handler {
	return &Handler{
		requestUsc: requestUsc,
	}
}

func (h *Handler) Init(api *echo.Group) {
	v1 := api.Group("/v1")
	{
		h.initRequestRoutes(v1)
	}
}
