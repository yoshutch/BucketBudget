package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"net/http"
	"os"
	"text/template"

	"github.com/yoshutch/BucketBudget/internal/data"
	"github.com/yoshutch/BucketBudget/internal/services"

	_ "github.com/lib/pq"
)

type application struct {
	logger        *slog.Logger
	templateCache map[string]*template.Template
	bucketService services.BucketsServiceI
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", os.Getenv("BUCKET_DB_DSN"), "PostgreSQL DSN")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	db, err := openDB(*dsn)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	logger.Info("connected to database successfully")
	defer db.Close()

	templateCache, err := newTemplateCache()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	bucketsRepo := data.BucketsRepo{DB: db}

	buckets := services.BucketsService{Logger: logger, Repo: &bucketsRepo}

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

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
