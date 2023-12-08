package routes

import (
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

func (s *service) NewRoutes() http.Handler {
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

	router.Route("/user", func(r chi.Router) {
		r.Post("/create", s.userService.CretaeUser)
		r.Post("/delete", s.userService.DeleteUser)
	})

	router.Route("/balance", func(r chi.Router) {
		r.Get("/info", s.balanceService.BalanceInfo)
		r.Post("/replenishment", s.balanceService.BalanceReplenishment)
		r.Patch("/debit", s.balanceService.BalanceDebit)
		r.Patch("/user-to-user", s.balanceService.UserToUser)
	})

	router.Route("/descriptions", func(r chi.Router) {
		r.Post("/add", s.descriptionService.AddDescription)
		r.Get("/get", s.descriptionService.GetDescriptions)
	})

	return router
}

func New(
	userService UserService,
	balanceService BalanceService,
	descriptionService DescriptionService,
) *service {
	return &service{
		userService:        userService,
		balanceService:     balanceService,
		descriptionService: descriptionService,
	}
}

type service struct {
	userService        UserService
	balanceService     BalanceService
	descriptionService DescriptionService
}
