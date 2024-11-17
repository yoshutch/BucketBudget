package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
	"text/template"

	"yosbomb.com/bucketbudget/internal/services"
)

type application struct {
	logger        *slog.Logger
	templateCache map[string]*template.Template
	bucketService services.BucketsServiceI
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	templateCache, err := newTemplateCache()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	buckets := services.NewBucketsService(logger)

	app := &application{
		logger:        logger,
		templateCache: templateCache,
		bucketService: &buckets,
	}

	srv := &http.Server{
		Addr:     *addr,
		Handler:  app.routes(),
		ErrorLog: slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	logger.Info("starting server", "addr", srv.Addr)
	err = srv.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)
}
