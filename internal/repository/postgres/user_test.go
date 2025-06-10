package postgres

import (
	"context"
	"lizobly/cotc-db/pkg/helpers"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_(t *testing.T) {
	db, mock, err := helpers.NewMockDB()

	if err != nil {
		t.Fatal(err)
	}
	repo := NewUserRepository(db)
	username := "username"
	t.Run("success", func(t *testing.T) {
		mock.ExpectQuery("SELECT (.*)").WithArgs(username, 1).WillReturnRows(
			sqlmock.NewRows([]string{"id"}).AddRow(1))
		res, err := repo.GetByUsername(context.TODO(), username)
		assert.NoError(t, err)
		assert.NotNil(t, res)
	})

}
