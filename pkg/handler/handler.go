package handler

import (
	"net/http"
	"tech_task/pkg/service"

	"github.com/go-chi/chi"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() http.Handler {
	router := chi.NewRouter()

	router.Route("/", func(r chi.Router) {
		r.Post("/up-balance", h.UpBalance)
		r.Patch("/writing-off", h.WritingOff)
		r.Get("/balance-info", h.BalanceInfo)
		r.Patch("/user-to-user", h.U2U)
	})

	router.Route("/description", func(r chi.Router) {
		r.Post("/add", h.AddDescription)
		r.Get("/get", h.GetDescriptions)
	})

	return router
}
