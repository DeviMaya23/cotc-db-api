package rest

import (
	"lizobly/cotc-db/domain"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TravellerService interface {
	GetByID(ctx echo.Context, id int) (res domain.Traveller, err error)
	Create(ctx echo.Context, input *domain.Traveller) (err error)
	Update(ctx echo.Context, input *domain.Traveller) (err error)
	Delete(ctx echo.Context, id int) (err error)
}

type TravellerHandler struct {
	Service TravellerService
}

func NewTravellerHandler(e *echo.Echo, svc TravellerService) {
	handler := &TravellerHandler{
		Service: svc,
	}
	e.GET("/travellers/:id", handler.GetByID)
	e.POST("/travellers", handler.Create)
	e.PUT("/travellers/:id", handler.Update)
	e.DELETE("/travellers/:id", handler.Delete)
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

func (a *TravellerHandler) Create(c echo.Context) error {

	var newTraveller domain.Traveller
	err := c.Bind(&newTraveller)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = a.Service.Create(c, &newTraveller)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, newTraveller)
}

func (a *TravellerHandler) Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "id not found")
	}

	var traveller domain.Traveller
	err = c.Bind(&traveller)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	traveller.ID = int64(id)

	err = a.Service.Update(c, &traveller)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, traveller)
}

func (a *TravellerHandler) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "id not found")
	}

	err = a.Service.Delete(c, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, nil)
}
