package tables

import (
	"database/sql"

	_ "github.com/glebarez/go-sqlite"
)

type SnippetModel struct {
	db *sql.DB
}

func CreateTable(db *sql.DB) (sql.Result, error) {
	sql := `CREATE TABLE IF NOT EXISTS countries (
        id INTEGER PRIMARY KEY,
        name     TEXT NOT NULL,
        population INTEGER NOT NULL,
        area INTEGER NOT NULL
    );`

	return db.Exec(sql)
}
