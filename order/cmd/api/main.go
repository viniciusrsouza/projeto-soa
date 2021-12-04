package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"github.com/viniciusrsouza/projeto-soa/order/config"
	"github.com/viniciusrsouza/projeto-soa/order/domain/usecases"
	"github.com/viniciusrsouza/projeto-soa/order/gateways"
	"github.com/viniciusrsouza/projeto-soa/order/gateways/storage"
	"github.com/viniciusrsouza/projeto-soa/order/gateways/storage/postgres"
)

const dbURL = "postgres://postgres:postgres@order_pg_db:5432/postgres?sslmode=disable"

func main() {
	logEntry := logrus.NewEntry(logrus.New())

	cfg, err := config.Load()
	if err != nil {
		logEntry.WithError(err).Fatal("could not load config")
	}

	logEntry = logEntry.WithFields(logrus.Fields{
		"app_name":    cfg.AppName,
		"port":        cfg.Port,
		"host":        cfg.Host,
		"environment": cfg.Environment,
	})

	ctx := context.Background()

	pool, err := pgxpool.Connect(ctx, dbURL)
	if err != nil {
		logEntry.WithError(err).Fatal("could not connect to postgres")
	}
	defer pool.Close()

	err = postgres.RunMigrations(dbURL)
	if err != nil {
		logEntry.WithError(err).Fatal("could not run migrations")
	}

	logEntry.Info("postgres connected successfully")

	storage := storage.NewOrderStorage(pool)
	usecase := usecases.NewOrderUseCase(storage)

	// kafkaPublisher := events.NewKafkaPublisher(cfg.KafkaConfig)
	// err = kafkaPublisher.Start()
	// if err != nil {
	// 	logEntry.WithError(err).Fatal("could not start kafka producers")
	// }

	// producer := producers.New(&kafkaPublisher)

	fmt.Println(err)
	api := gateways.NewAPI(usecase, logEntry)

	h := api.BuildHandler()

	logEntry.Infof("starting api on %s:%s", cfg.Host, cfg.Port)
	logEntry.Fatal(http.ListenAndServe(":"+cfg.Port, h))
}
