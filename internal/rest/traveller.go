package rest

import (
	"lizobly/cotc-db/domain"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TravellerService interface {
	GetByID(ctx echo.Context, id int) (res domain.Traveller, err error)
}

type TravellerHandler struct {
	Service TravellerService
}

func NewTravellerHandler(e *echo.Echo, svc TravellerService) {
	handler := &TravellerHandler{
		Service: svc,
	}
	e.GET("/travellers/:id", handler.GetByID)
}

func (a *TravellerHandler) GetByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "id not found")
	}

	traveller, err := a.Service.GetByID(c, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, traveller)
}
