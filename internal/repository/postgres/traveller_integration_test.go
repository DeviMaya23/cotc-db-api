package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"lizobly/cotc-db/pkg/domain"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"

	pgGormDriver "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestTravellerRepository_Integration(t *testing.T) {
	ctx := context.Background()

	tespath := filepath.Join("../../..", "testdata", "db-traveller-repo.sql")
	fmt.Println(tespath)

	pgContainer, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("postgres:15.3-alpine"),
		postgres.WithInitScripts(tespath),
		postgres.WithDatabase("test-db"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		if err := pgContainer.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate pgContainer: %s", err)
		}
	})

	connStr, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
	assert.NoError(t, err)

	dbConn, err := sql.Open("pgx", connStr)
	if err != nil {
		t.Fatal("failed open database ", err)
	}
	db, err := gorm.Open(pgGormDriver.New(pgGormDriver.Config{
		Conn: dbConn,
	}), &gorm.Config{})
	if err != nil {
		t.Fatal("failed to open gorm ", err)
	}

	repo := NewTravellerRepository(db)

	errCreate := repo.Create(context.TODO(), &domain.Traveller{
		Name:        "Fiore",
		Rarity:      5,
		InfluenceID: 3,
	})
	assert.Nil(t, errCreate)

	traveller, err := repo.GetByID(context.TODO(), 1)
	assert.Nil(t, err)
	assert.Equal(t, traveller.Name, "Fiore")
	assert.Equal(t, traveller.Rarity, 5)
	assert.Equal(t, traveller.InfluenceID, 3)

	// Update traveller
	err = repo.Update(context.TODO(), &domain.Traveller{
		CommonModel: domain.CommonModel{
			ID: 1,
		},
		Rarity: 6,
	})
	assert.Nil(t, err)

	// Check updated traveller
	traveller, err = repo.GetByID(context.TODO(), 1)
	assert.Nil(t, err)
	assert.Equal(t, traveller.Name, "Fiore")
	assert.Equal(t, traveller.Rarity, 6)

	// Delete traveller
	err = repo.Delete(context.TODO(), 1)
	assert.Nil(t, err)
}
