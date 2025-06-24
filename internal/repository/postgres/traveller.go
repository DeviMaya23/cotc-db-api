package postgres

import (
	"context"
	"lizobly/cotc-db-api/pkg/domain"

	"gorm.io/gorm"
)

type TravellerRepository struct {
	db *gorm.DB
}

func NewTravellerRepository(db *gorm.DB) *TravellerRepository {
	return &TravellerRepository{
		db: db,
	}
}
func (r TravellerRepository) GetByID(ctx context.Context, id int) (result domain.Traveller, err error) {
	err = r.db.WithContext(ctx).Preload("Influence").First(&result, "id = ?", id).Error
	return
}

func (r TravellerRepository) Create(ctx context.Context, input *domain.Traveller) (err error) {
	err = r.db.WithContext(ctx).Create(input).Error
	return
}

func (r TravellerRepository) Update(ctx context.Context, input *domain.Traveller) (err error) {
	err = r.db.WithContext(ctx).Updates(input).Error
	return
}

func (r TravellerRepository) Delete(ctx context.Context, id int) (err error) {
	err = r.db.WithContext(ctx).Delete(&domain.Traveller{}, id).Error
	return
}
