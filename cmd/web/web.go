package main

import (
	"flag"
	"fmt"
	"grpc-web-template/internal/common"
	"grpc-web-template/internal/repository"
	"grpc-web-template/internal/repository/sqliteRepo"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"

	"github.com/alexedwards/scs/sqlite3store"
	"github.com/alexedwards/scs/v2"
)

var session *scs.SessionManager
var version = "1.0"

type config struct {
	port       int
	backend    string
	apiService string
	apiProxy   string
	sqlitePath string
}

type application struct {
	config        config
	infoLog       *log.Logger
	errorLog      *log.Logger
	templateCache map[string]*template.Template
	DB            repository.Repository
	Session       *scs.SessionManager
	version       string
}

func main() {
	var cfg config

	flag.StringVar(&cfg.apiService, "apiService", ":9093", "URL to grpc service")
	flag.StringVar(&cfg.backend, "backend", ":3500", "URL to api")
	flag.StringVar(&cfg.apiProxy, "apiProxy", ":8082", "proxy to 9093 service")
	flag.StringVar(&cfg.sqlitePath, "sqlitePath", "./sql.db", "url to sqlite database")
	flag.IntVar(&cfg.port, "port", 3600, "Server port to listen on")

	flag.StringVar(&common.TimeFormat, "timeFormat", "2006-01-02 15:04:05.000", "timeFormat")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// connect to sqlite database
	conn, err := sqliteRepo.ConnectSQL(cfg.sqlitePath)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// set up session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Store = sqlite3store.NewWithCleanupInterval(conn, 5*time.Minute)

	tc := make(map[string]*template.Template)

	// create application
	app := &application{
		config:        cfg,
		infoLog:       infoLog,
		errorLog:      errorLog,
		templateCache: tc,
		Session:       session,
		DB:            sqliteRepo.SetupDB(conn),
		version:       version,
	}

	err = app.serve()
	if err != nil {
		app.errorLog.Println(err)
		log.Fatal(err)
	}
}

func (app *application) serve() error {
	// serve grpc web api service
	go app.startGrpcApiService()

	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", app.config.port),
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	app.infoLog.Printf("Starting HTTP server on port %d\n", app.config.port)

	return srv.ListenAndServe()
}
