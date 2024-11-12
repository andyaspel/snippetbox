package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/andyaspel/snippetbox/pkg/models"
	"github.com/andyaspel/snippetbox/pkg/models/sqlte"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	snippets *sqlte.SnippetModel
}

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "\nINFO:\n\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "\nERROR:\n\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := connectToSQLite()
	if err != nil {
		errorLog.Fatal(err)
	}
	var Snippet models.Snippet
	err = db.AutoMigrate(&Snippet)
	if err != nil {
		log.Fatal(err)
	}

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		snippets: &sqlte.SnippetModel{DB: db},
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)

}

func connectToSQLite() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("./db/snippets.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
