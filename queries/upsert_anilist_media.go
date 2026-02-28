package queries

import (
	"context"
	"database/sql"
	_ "embed"
)

//go:embed upsert_anilist_media.sql
var upsertAnilistMediaSQL string

type UpsertAnilistMediaParameters struct {
	Chapters sql.NullInt64
	Episodes sql.NullInt64
	Format   string
	ID       int64
	Title    string
	Volumes  sql.NullInt64
}

func (q *Queries) UpsertAnilistMedia(c context.Context, p UpsertAnilistMediaParameters) error {
	_, err := q.connection.ExecContext(c, upsertAnilistMediaSQL,
		p.Chapters,
		p.Episodes,
		p.Format,
		p.ID,
		p.Title,
		p.Volumes,
	)
	return err
}
