package postgres

import (
	"context"
	"lizobly/cotc-db/pkg/domain"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}
func (r UserRepository) GetByUsername(ctx context.Context, username string) (result domain.User, err error) {
	err = r.db.WithContext(ctx).First(&result, "username = ?", username).Error
	return
}
