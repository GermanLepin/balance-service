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

	"tech_task/internal/godb"
	"tech_task/internal/handler"
)

func main() {
	ctx := context.Background()
	r := chi.NewRouter()

	service := handler.HttpService{
		UserService: &godb.Instance{},
	}

	r.Route("/", func(r chi.Router) {
		r.Post("/up-balance", service.UpBalance)
		r.Patch("/writing-off", service.WritingOff)
		r.Get("/balance-info", service.BalanceInfo)
		r.Patch("/user-to-user", service.U2U)
	})

	r.Route("/description", func(r chi.Router) {
		r.Post("/add", service.AddDescription)
		r.Get("/get-user", service.GetUserId)
		r.Get("/get-user/sort_by", service.GetUserIdSort)
		r.Get("/get-user/sort_by/desc", service.GetUserIdSortDesc)
		r.Get("/get-all", service.GetAllUsers)
		r.Get("/get-all/sort_by", service.GetAllUsersSort)
		r.Get("/get-all/sort_by/desc", service.GetAllUsersSortDesc)
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
