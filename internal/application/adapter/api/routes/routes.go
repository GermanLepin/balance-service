package routes

import (
	"balance-service/internal/application/adapter/api/http/create_user_handler"
	"balance-service/internal/application/repository"
	"balance-service/internal/application/service/create_user_service"
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type (
	UserService interface {
		CretaeUser(w http.ResponseWriter, r *http.Request)
		DeleteUser(w http.ResponseWriter, r *http.Request)
	}

	BalanceService interface {
		BalanceInfo(w http.ResponseWriter, r *http.Request)
		BalanceReplenishment(w http.ResponseWriter, r *http.Request)
		BalanceDebit(w http.ResponseWriter, r *http.Request)
		UserToUser(w http.ResponseWriter, r *http.Request)
	}

	DescriptionService interface {
		AddDescription(w http.ResponseWriter, r *http.Request)
		GetDescriptions(w http.ResponseWriter, r *http.Request)
	}
)

func (s *service) NewRoutes(connection *sql.DB) http.Handler {
	router := chi.NewRouter()

	// specify who is allowed to connect
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	userRepository := repository.NewUserRepository(connection)
	cretae_user_service := create_user_service.New(userRepository)
	create_user_handler := create_user_handler.New(cretae_user_service)

	router.Route("/user", func(r chi.Router) {
		r.Post("/create", create_user_handler.CretaeUser)
		//	r.Post("/delete", s.userService.DeleteUser)
	})

	router.Route("/balance", func(r chi.Router) {
		// r.Get("/info", s.balanceService.BalanceInfo)
		// r.Post("/replenishment", s.balanceService.BalanceReplenishment)
		// r.Patch("/debit", s.balanceService.BalanceDebit)
		// r.Patch("/user-to-user", s.balanceService.UserToUser)
	})

	router.Route("/descriptions", func(r chi.Router) {
		//r.Post("/add", s.descriptionService.AddDescription)
		// r.Get("/get", s.descriptionService.GetDescriptions)
	})

	return router
}

func New(

// userService UserService,
// balanceService BalanceService,
// descriptionService DescriptionService,
) *service {
	return &service{
		// userService:        userService,
		// balanceService:     balanceService,
		// descriptionService: descriptionService,
	}
}

type service struct {
	// userService        UserService
	// balanceService     BalanceService
	// descriptionService DescriptionService
}
