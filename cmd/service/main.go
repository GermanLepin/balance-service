package main

import (
	"fmt"
	"log"
	"net/http"

	"balance-service/db/postgres/connection"
	"balance-service/internal/application/adapter/api/http/create_user_handler"
	"balance-service/internal/application/adapter/api/http/delete_user_handler"
	"balance-service/internal/application/adapter/api/routes"
	"balance-service/internal/application/repository"
	"balance-service/internal/application/service/create_user_service"
	"balance-service/internal/application/service/delete_user_service"
	"balance-service/internal/application/service/json_service"
)

const (
	webPort = "80"

	accessKey = "27c4039d0e33e2f74fbdc7afa63c08a8"
)

func main() {
	connection := connection.StartDB()
	jsonService := json_service.New()

	userRepository := repository.NewUserRepository(connection)

	delete_user_service := delete_user_service.New(userRepository)
	cretae_user_service := create_user_service.New(userRepository)

	create_user_handler := create_user_handler.New(cretae_user_service, jsonService)
	delete_user_handler := delete_user_handler.New(delete_user_service, jsonService)

	api_routes := routes.New(
		connection,
		create_user_handler,
		delete_user_handler,
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
