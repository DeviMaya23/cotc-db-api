package traveller

import (
	"context"
	"lizobly/cotc-db/pkg/constants"
	"lizobly/cotc-db/pkg/domain"

	"github.com/labstack/echo/v4"
)

type TravellerRepository interface {
	GetByID(ctx context.Context, id int) (result domain.Traveller, err error)
	Create(ctx context.Context, input *domain.Traveller) (err error)
	Update(ctx context.Context, input *domain.Traveller) (err error)
	Delete(ctx context.Context, id int) (err error)
}

type Service struct {
	travellerRepo TravellerRepository
}

func NewTravellerService(t TravellerRepository) *Service {
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

func (s Service) Create(ctx echo.Context, input domain.CreateTravellerRequest) (err error) {

	newTraveller := domain.Traveller{
		Name:        input.Name,
		Rarity:      input.Rarity,
		InfluenceID: constants.GetInfluenceID(input.Influence),
	}

	err = s.travellerRepo.Create(ctx.Request().Context(), &newTraveller)
	if err != nil {
		// TODO: log
		return
	}
	return
}

func (s Service) Update(ctx echo.Context, input *domain.Traveller) (err error) {

	err = s.travellerRepo.Update(ctx.Request().Context(), input)
	if err != nil {
		// TODO: log
		return
	}
	return
}

func (s Service) Delete(ctx echo.Context, id int) (err error) {

	err = s.travellerRepo.Delete(ctx.Request().Context(), id)
	if err != nil {
		// TODO: log
		return
	}
	return
}
