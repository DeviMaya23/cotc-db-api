package main

import (
	"database/sql"
	"fmt"
	postgresRepo "lizobly/cotc-db/internal/repository/postgres"
	"lizobly/cotc-db/internal/rest"
	"lizobly/cotc-db/pkg/validator"
	"lizobly/cotc-db/traveller"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbUser := os.Getenv("DATABASE_USER")
	dbPass := os.Getenv("DATABASE_PASS")
	dbName := os.Getenv("DATABASE_NAME")
	dsn := fmt.Sprintf("sslmode=disable host=%s port=%s user=%s password='%s' dbname=%s timezone=%s", dbHost, dbPort, dbUser, dbPass, dbName, "Asia/Jakarta")

	dbConn, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal("failed open database ", err)
	}
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: dbConn,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to open gorm ", err)
	}

	err = dbConn.Ping()
	if err != nil {
		log.Fatal("failed to ping database ", err)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal("got error when closing the DB connection", err)
		}
	}()

	addr := fmt.Sprintf(":%s", os.Getenv("APP_PORT"))
	e := echo.New()

	// Validator
	e.Validator = validator.NewValidator()

	// Repository
	travellerRepo := postgresRepo.NewTravellerRepository(db)

	// Service
	travellerService := traveller.NewService(travellerRepo)

	// Handler
	rest.NewTravellerHandler(e, travellerService)

	e.Logger.Fatal(e.Start(addr))
}
