package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"avito_tech_task/db/postgres/connection"
	"avito_tech_task/internal/application/adapter/api/routes"
)

const webPort = "80"

func main() {
	connection := startDB()

	jsonService := json_service.New()

	userService := userService.New()
	balanceService := balanceService.New()
	descriptionService := descriptionService.New()

	api_routes := routes.New(userService, balanceService, descriptionService)

	log.Printf("starting avito tech task service on port %s\n", webPort)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: api_routes.NewRoutes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

func startDB() *sql.DB {
	cfg := &connection.Config{}

	if len(os.Getenv("POSTGRES_HOST")) == 0 {
		cfg.Host = "localhost"
	} else {
		cfg.Host = os.Getenv("POSTGRES_HOST")
	}

	if len(os.Getenv("POSTGRES_USER")) == 0 {
		cfg.Username = "postgres"
	} else {
		cfg.Username = os.Getenv("POSTGRES_USER")
	}

	if len(os.Getenv("POSTGRES_PASSWORD")) == 0 {
		cfg.Password = "password"
	} else {
		cfg.Password = os.Getenv("POSTGRES_PASSWORD")
	}

	if len(os.Getenv("POSTGRES_DB")) == 0 {
		cfg.DbName = "avito_db"
	} else {
		cfg.DbName = os.Getenv("POSTGRES_DB")
	}

	cfg.Port = "54320"
	cfg.Timeout = 5

	return connection.NewConnection(cfg)
}
