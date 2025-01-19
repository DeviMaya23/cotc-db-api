package main

import (
	"database/sql"
	"fmt"
	_ "lizobly/cotc-db/docs"
	"lizobly/cotc-db/user"

	postgresRepo "lizobly/cotc-db/internal/repository/postgres"
	"lizobly/cotc-db/internal/rest"
	"lizobly/cotc-db/pkg/middleware"
	"lizobly/cotc-db/pkg/validator"
	"lizobly/cotc-db/traveller"
	"log"
	"os"

	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//	@title			COTC DB API
//	@version		1.0
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	Liz
//	@contact.email	j2qgehn84@mozmail.com

// @BasePath	/api/v1
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

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	// Validator
	validator := validator.NewValidator()
	e.Validator = validator
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			ctx.Set("validator", validator)
			return next(ctx)
		}
	})

	// Middleware
	jwtMiddleware := middleware.NewJWTMiddleware()

	// Repository
	travellerRepo := postgresRepo.NewTravellerRepository(db)
	userRepo := postgresRepo.NewUserRepository(db)

	// Service
	travellerService := traveller.NewTravellerService(travellerRepo)
	userService := user.NewUserService(userRepo)

	v1 := e.Group("/api/v1")
	v1.Use(jwtMiddleware)

	// Handler
	rest.NewTravellerHandler(v1, travellerService)

	// NewUserHandler
	echoGroup := e.Group("/")
	rest.NewUserHandler(echoGroup, userService)

	e.Logger.Fatal(e.Start(addr))
}
