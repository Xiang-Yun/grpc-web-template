package main

import (
	"flag"
	"fmt"
	"grpc-web-template/internal/common"
	"grpc-web-template/internal/repository"
	"grpc-web-template/internal/repository/sqliteRepo"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type config struct {
	port       int
	sqlitePath string
}

type application struct {
	infoLog  *log.Logger
	errorLog *log.Logger
	config   config
	DB       repository.Repository
}

func main() {
	var cfg config
	flag.IntVar(&cfg.port, "port", 3500, "Server port to listen on")
	flag.StringVar(&cfg.sqlitePath, "sqlitePath", "./sql.db", "url to sqlite database")
	flag.StringVar(&common.TimeFormat, "timeFormat", "2006-01-02 15:04:05.000", "timeFormat")
	flag.Parse()

	// create info and error logger
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime)

	// connect to sqlite database
	conn, err := sqliteRepo.ConnectSQL(cfg.sqlitePath)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	app := &application{
		infoLog:  infoLog,
		errorLog: errorLog,
		config:   cfg,
		DB:       sqliteRepo.SetupDB(conn),
	}

	err = app.serve()
	if err != nil {
		app.errorLog.Println(err)
		log.Fatal(err)
	}
}

func (app *application) serve() error {
	src := &http.Server{
		Addr:              fmt.Sprintf(":%d", app.config.port),
		Handler:           app.routes(),
		IdleTimeout:       60 * time.Second,
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
	}
	app.infoLog.Printf("Starting Back end server  on port %d\n", app.config.port)
	return src.ListenAndServe()
}
