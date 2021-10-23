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
		r.Get("/balance-info", handlers.BalanceInfo)
		r.Patch("/writing-off", handlers.WritingOff)
		r.Patch("/user-to-user", handlers.U2U)
		r.Delete("/delete-user", handlers.DeleteUser)
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
			log.WithError(err).Fatal("start server")
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
