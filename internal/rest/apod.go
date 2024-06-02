package rest

import (
	"NasaEnjoyer/domain"
	"context"

	"github.com/labstack/echo/v4"
)

type ResponceError struct {
	Message string `json:"message"`
}

type AdopService interface {
	GetAPODs(ctx context.Context) ([]domain.APOD, error)
	GetAPOD(ctx context.Context, date string) (*domain.APOD, error)
	FetchAndStoreAPOD(ctx context.Context) error
}

type APODHandler struct {
	apodService AdopService
}

func NewAPODHandler(apodService AdopService, e *echo.Echo) {
	handler := &APODHandler{
		apodService: apodService,
	}

	e.GET("/apod", handler.GetAPODs)

	e.GET("/apod/:date", handler.GetAPOD)
	e.GET("/apod/fetch", handler.FetchAndStoreAPOD)
}

func (h *APODHandler) GetAPODs(c echo.Context) error {
	ctx := c.Request().Context()

	apods, err := h.apodService.GetAPODs(ctx)
	if err != nil {
		return c.JSON(500, ResponceError{Message: err.Error()})
	}

	return c.JSON(200, apods)
}

func (h *APODHandler) GetAPOD(c echo.Context) error {
	ctx := c.Request().Context()
	date := c.Param("date")

	apod, err := h.apodService.GetAPOD(ctx, date)
	if err != nil {
		return c.JSON(500, ResponceError{Message: err.Error()})
	}

	return c.JSON(200, apod)
}

func (h *APODHandler) FetchAndStoreAPOD(c echo.Context) error {
	ctx := c.Request().Context()

	err := h.apodService.FetchAndStoreAPOD(ctx)
	if err != nil {
		return c.JSON(500, ResponceError{Message: err.Error()})
	}

	return c.JSON(200, nil)
}
