package main

import (
	"fmt"
	"log"
	"net/http"

	"balance-service/db/postgres/connection"
	"balance-service/internal/application/adapter/api/http/create_user_handler"
	"balance-service/internal/application/adapter/api/http/delete_user_handler"
	"balance-service/internal/application/adapter/api/http/fetch_balance_info_handler"
	"balance-service/internal/application/adapter/api/routes"
	"balance-service/internal/application/repository"
	"balance-service/internal/application/service/create_user_service"
	"balance-service/internal/application/service/delete_user_service"
	"balance-service/internal/application/service/fetch_balance_info_service"
	"balance-service/internal/application/service/json_service"

	"github.com/joho/godotenv"
)

const (
	webPort = "80"

	accessKey = "27c4039d0e33e2f74fbdc7afa63c08a8"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("error loading .env file")
	}
}

func main() {
	connection := connection.StartDB()
	jsonService := json_service.New()

	userRepository := repository.NewUserRepository(connection)

	cretae_user_service := create_user_service.New(userRepository)
	delete_user_service := delete_user_service.New(userRepository)
	fetch_balance_info_service := fetch_balance_info_service.New(userRepository)

	create_user_handler := create_user_handler.New(cretae_user_service, jsonService)
	delete_user_handler := delete_user_handler.New(delete_user_service, jsonService)
	fetch_balance_info_handler := fetch_balance_info_handler.New(fetch_balance_info_service, jsonService)

	api_routes := routes.New(
		connection,
		create_user_handler,
		delete_user_handler,
		fetch_balance_info_handler,
	)

	log.Printf("starting balance service on port %s\n", webPort)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: api_routes.NewRoutes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
