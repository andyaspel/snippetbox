package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/glebarez/go-sqlite"
)

func main() {
	// var db  *sql.DB
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()
	// snippets := db
	infoLog := log.New(os.Stdout, "\nINFO:\n\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "\nERROR:\n\t", log.Ldate|log.Ltime|log.Lshortfile)
	// var snippets *sql.DB
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)

	// Connect to the SQLite database
	db, err := sql.Open(
		"sqlite",
		"db/snippets.db?_pragma=foreign_keys(1)",
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()
}
