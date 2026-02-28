package queries

import (
	"context"
	"database/sql"
	_ "embed"
)

//go:embed list_anime.sql
var listAnimeSQL string

type ListAnimeRow struct {
	Favorite        int64
	ProgressCurrent int64
	ProgressTotal   sql.NullInt64
	Repeat          int64
	Title           string
}

func (q *Queries) ListAnime(c context.Context) ([]ListAnimeRow, error) {
	rows, err := q.connection.QueryContext(c, listAnimeSQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListAnimeRow
	for rows.Next() {
		var row ListAnimeRow
		if err := rows.Scan(
			&row.Favorite,
			&row.ProgressCurrent,
			&row.ProgressTotal,
			&row.Repeat,
			&row.Title,
		); err != nil {
			return nil, err
		}
		items = append(items, row)
	}
	return items, rows.Err()
}
