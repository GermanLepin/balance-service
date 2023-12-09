package main

import (
	"fmt"
	"log"
	"net/http"

	"balance-service/db/postgres/connection"
	"balance-service/internal/application/adapter/api/routes"
)

const webPort = "8011"

func main() {
	connection := connection.StartDB()

	// jsonService := json_service.New()

	// userService := userService.New()
	// balanceService := balanceService.New()
	// descriptionService := descriptionService.New()

	// api_routes := routes.New(userService, balanceService, descriptionService)

	api_routes := routes.New()

	log.Printf("starting balance service on port %s\n", webPort)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: api_routes.NewRoutes(connection),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
