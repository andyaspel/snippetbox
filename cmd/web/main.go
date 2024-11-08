package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/andyaspel/snippetbox/pkg/tables"
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
		snippets: tables.CreateTable,
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

	_, err = app.CreateTable(db)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Connected to the SQLite database successfully.")

	fmt.Println("Table countries was created successfully.")

	// Get the version of SQLite
	var sqliteVersion string
	err = db.QueryRow("select sqlite_version()").Scan(&sqliteVersion)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(sqliteVersion)
}
