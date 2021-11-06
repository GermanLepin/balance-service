package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"

	log "github.com/sirupsen/logrus"

	"tech_task/internal/handlers/balance"
	"tech_task/internal/handlers/description"
	topUp "tech_task/internal/handlers/top_up"
	"tech_task/internal/handlers/u2u"
	writingOff "tech_task/internal/handlers/writing_off"
)

func main() {
	ctx := context.Background()
	r := chi.NewRouter()

	r.Route("/", func(r chi.Router) {
		r.Post("/up-balance", topUp.UpBalance)
		r.Patch("/writing-off", writingOff.WritingOff)
		r.Get("/balance-info", balance.BalanceInfo)
		r.Patch("/user-to-user", u2u.U2U)
	})

	r.Route("/description", func(r chi.Router) {
		r.Post("/add", description.AddDescription)
		r.Get("/get-user", description.GetUserId)
		r.Get("/get-user/sort_by", description.GetUserIdSort)
		r.Get("/get-user/sort_by/desc", description.GetUserIdSortDesc)
		r.Get("/get-all", description.GetAllUsers)
		r.Get("/get-all/sort_by", description.GetAllUsersSort)
		r.Get("/get-all/sort_by/desc", description.GetAllUsersSortDesc)
	})

	srv := &http.Server{
		Addr:    ":9000",
		Handler: r,
	}
	defer srv.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.WithError(err).Fatal("Error start server")
		}
	}()

	<-stop
	log.Info("caught stop signal")

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.WithError(err).Fatal("Server Shutdown Failed")
	}

	log.Info("Server Exited Properly")
}
