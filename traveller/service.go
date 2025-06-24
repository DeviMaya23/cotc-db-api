package traveller

import (
	"context"
	"lizobly/cotc-db-api/pkg/constants"
	"lizobly/cotc-db-api/pkg/domain"
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

func (s Service) GetByID(ctx context.Context, id int) (res domain.Traveller, err error) {

	res, err = s.travellerRepo.GetByID(ctx, id)
	if err != nil {
		// TODO: log
		return
	}
	return
}

func (s Service) Create(ctx context.Context, input domain.CreateTravellerRequest) (err error) {

	newTraveller := domain.Traveller{
		Name:        input.Name,
		Rarity:      input.Rarity,
		InfluenceID: constants.GetInfluenceID(input.Influence),
	}

	err = s.travellerRepo.Create(ctx, &newTraveller)
	if err != nil {
		// TODO: log
		return
	}
	return
}

func (s Service) Update(ctx context.Context, input *domain.Traveller) (err error) {

	err = s.travellerRepo.Update(ctx, input)
	if err != nil {
		// TODO: log
		return
	}
	return
}

func (s Service) Delete(ctx context.Context, id int) (err error) {

	err = s.travellerRepo.Delete(ctx, id)
	if err != nil {
		// TODO: log
		return
	}
	return
}
