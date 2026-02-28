package main

import (
	"anictl/queries"
	"database/sql"
	_ "embed"
	"strings"

	_ "modernc.org/sqlite"
)

//go:embed schema.sql
var schema string

func openDatabase() (*queries.Queries, *sql.DB, error) {
	database, err := sql.Open("sqlite", "database.sqlite")
	if err != nil {
		return nil, nil, err
	}
	for statement := range strings.SplitSeq(schema, ";") {
		statement = strings.TrimSpace(statement)
		if statement == "" {
			continue
		}
		if _, err := database.Exec(statement); err != nil {
			database.Close()
			return nil, nil, err
		}
	}
	return queries.New(database), database, nil
}
