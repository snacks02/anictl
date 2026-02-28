package queries

import (
	"context"
	"database/sql"
)

type dbtx interface {
	ExecContext(context.Context, string, ...any) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...any) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...any) *sql.Row
}

type Queries struct {
	connection dbtx
}

func New(database *sql.DB) *Queries {
	return &Queries{connection: database}
}

func (q *Queries) WithTx(transaction *sql.Tx) *Queries {
	return &Queries{connection: transaction}
}
