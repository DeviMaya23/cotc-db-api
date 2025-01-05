package rest

import (
	"lizobly/cotc-db/pkg/domain"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TravellerService interface {
	GetByID(ctx echo.Context, id int) (res domain.Traveller, err error)
	Create(ctx echo.Context, input domain.CreateTravellerRequest) (err error)
	Update(ctx echo.Context, input *domain.Traveller) (err error)
	Delete(ctx echo.Context, id int) (err error)
}

type TravellerHandler struct {
	Controller
	Service TravellerService
}

func NewTravellerHandler(e *echo.Echo, svc TravellerService) {
	handler := &TravellerHandler{
		Service: svc,
	}
	v1 := e.Group("/api/v1/travellers")

	v1.GET("/:id", handler.GetByID)
	v1.POST("", handler.Create)
	v1.PUT("/:id", handler.Update)
	v1.DELETE("/:id", handler.Delete)
}

// GetByID godoc
//
//	@Summary		Get by ID
//	@Description	get traveller information by ID
//	@Tags			accounts
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Account ID"
//	@Success		200	{object}	domain.Traveller
//	@Failure		400	{object}	StandardAPIResponse
//	@Failure		404	{object}	StandardAPIResponse
//	@Failure		500	{object}	StandardAPIResponse
//	@Router			/travellers/{id} [get]
func (a *TravellerHandler) GetByID(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return a.ResponseError(ctx, http.StatusBadRequest, "error validation", "id not found")
	}

	traveller, err := a.Service.GetByID(ctx, id)
	if err != nil {
		return a.ResponseError(ctx, http.StatusBadRequest, "error get data", err.Error())
	}

	return a.Ok(ctx, "success", traveller, nil)
}

func (a *TravellerHandler) Create(ctx echo.Context) error {

	var newTraveller domain.CreateTravellerRequest
	err := ctx.Bind(&newTraveller)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	err = ctx.Validate(&newTraveller)
	if err != nil {
		return a.ResponseErrorValidation(ctx, err)
	}

	err = a.Service.Create(ctx, newTraveller)
	if err != nil {
		return a.ResponseError(ctx, http.StatusBadRequest, "error get data", err.Error())
	}

	return a.Ok(ctx, "success", newTraveller, nil)
}

func (a *TravellerHandler) Update(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return a.ResponseError(ctx, http.StatusBadRequest, "error validation", "id not found")
	}

	var traveller domain.Traveller
	err = ctx.Bind(&traveller)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	traveller.ID = int64(id)

	err = a.Service.Update(ctx, &traveller)
	if err != nil {
		return a.ResponseError(ctx, http.StatusBadRequest, "error get data", err.Error())
	}

	return a.Ok(ctx, "success", traveller, nil)
}

func (a *TravellerHandler) Delete(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return a.ResponseError(ctx, http.StatusBadRequest, "error validation", "id not found")
	}

	err = a.Service.Delete(ctx, id)
	if err != nil {
		return a.ResponseError(ctx, http.StatusBadRequest, "error get data", err.Error())
	}

	return a.Ok(ctx, "success", nil, nil)
}
