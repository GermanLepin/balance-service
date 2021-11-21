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

	"tech_task/internal/handlers"
)

func main() {
	ctx := context.Background()
	r := chi.NewRouter()

	r.Route("/", func(r chi.Router) {
		r.Post("/up-balance", handlers.UpBalance)
		r.Patch("/writing-off", handlers.WritingOff)
		r.Get("/balance-info", handlers.BalanceInfo)
		r.Patch("/user-to-user", handlers.U2U)
	})

	r.Route("/description", func(r chi.Router) {
		r.Post("/add", handlers.AddDescription)
		r.Get("/get-user", handlers.GetUserId)
		r.Get("/get-user/sort_by", handlers.GetUserIdSort)
		r.Get("/get-user/sort_by/desc", handlers.GetUserIdSortDesc)
		r.Get("/get-all", handlers.GetAllUsers)
		r.Get("/get-all/sort_by", handlers.GetAllUsersSort)
		r.Get("/get-all/sort_by/desc", handlers.GetAllUsersSortDesc)
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
