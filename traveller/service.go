package traveller

import (
	"context"
	"lizobly/cotc-db/domain"

	"github.com/labstack/echo/v4"
)

type TravellerRepository interface {
	GetByID(ctx context.Context, id int) (result domain.Traveller, err error)
}

type Service struct {
	travellerRepo TravellerRepository
}

func NewService(t TravellerRepository) *Service {
	return &Service{
		travellerRepo: t,
	}
}

func (s Service) GetByID(ctx echo.Context, id int) (res domain.Traveller, err error) {

	res, err = s.travellerRepo.GetByID(ctx.Request().Context(), id)
	if err != nil {
		// TODO: log
		return
	}

	return
}
