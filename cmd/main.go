package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"tech_task"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"tech_task/pkg/handler"
	"tech_task/pkg/repository"
	"tech_task/pkg/service"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initialixing configs: %v", err.Error())
	}

	db, err := StartDB()
	if err != nil {
		logrus.Fatalf("error connection database: %v", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(tech_task.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("Tech_task Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}
	logrus.Print("Tech_task Shutting Down")
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()

}

func StartDB() (*pgxpool.Pool, error) {
	cfg := &repository.Config{}

	if len(os.Getenv("POSTGRES_HOST")) == 0 {
		cfg.Host = "localhost"
	} else {
		cfg.Host = os.Getenv("POSTGRES_HOST")
	}

	if len(os.Getenv("POSTGRES_USER")) == 0 {
		cfg.Username = "postgres"
	} else {
		cfg.Username = os.Getenv("POSTGRES_USER")
	}

	if len(os.Getenv("POSTGRES_PASSWORD")) == 0 {
		cfg.Password = "1234"
	} else {
		cfg.Password = os.Getenv("POSTGRES_PASSWORD")
	}

	if len(os.Getenv("POSTGRES_DB")) == 0 {
		cfg.DbName = "avito_users_db"
	} else {
		cfg.DbName = os.Getenv("POSTGRES_DB")
	}

	cfg.Port = "54320"
	cfg.Timeout = 5

	poolConfig, err := repository.NewPoolConfig(cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Pool config error: %v\n", err)
		os.Exit(1)
	}

	poolConfig.MaxConns = 20

	c, err := repository.NewConnection(poolConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Connect to database failed: %v\n", err)
		os.Exit(1)
	}

	_, err = c.Exec(context.Background(), ";")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ping failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Ping OK!")

	return c, nil
}
