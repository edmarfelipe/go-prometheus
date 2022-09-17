package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/dlmiddlecote/sqlstats"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/edmarfelipe/go-prometheus/config"
	"github.com/edmarfelipe/go-prometheus/handlers"
	"github.com/edmarfelipe/go-prometheus/middlewares"
)

type App struct {
	router *mux.Router
	db     *sql.DB
}

func New(conf *config.Config) *App {
	app := &App{
		router: mux.NewRouter(),
		db:     createDBConnection(conf.DB),
	}

	return app.registerRoutes()
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.router))
}

func createDBConnection(conf *config.Database) *sql.DB {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", conf.Host, conf.Port, conf.User, conf.Password, conf.Name)
	log.Println(connectionString)

	dbConn, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	collector := sqlstats.NewStatsCollector(conf.Name, dbConn)
	prometheus.MustRegister(collector)

	return dbConn
}

func (a *App) registerRoutes() *App {
	a.router.Use(middlewares.PrometheusHTTPDuration)

	a.router.Path("/").
		Methods("GET").
		HandlerFunc(handlers.GetBooks(a.db))

	a.router.Path("/metrics").
		Methods("GET").
		Handler(promhttp.Handler())

	return a
}
