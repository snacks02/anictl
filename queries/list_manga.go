package queries

import (
	"context"
	"database/sql"
	_ "embed"
)

//go:embed list_manga.sql
var listMangaSQL string

type ListMangaRow struct {
	ChaptersCurrent int64
	ChaptersTotal   sql.NullInt64
	Favorite        int64
	Title           string
	VolumesCurrent  int64
	VolumesTotal    sql.NullInt64
}

func (q *Queries) ListManga(c context.Context) ([]ListMangaRow, error) {
	rows, err := q.connection.QueryContext(c, listMangaSQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListMangaRow
	for rows.Next() {
		var row ListMangaRow
		if err := rows.Scan(
			&row.ChaptersCurrent,
			&row.ChaptersTotal,
			&row.Favorite,
			&row.Title,
			&row.VolumesCurrent,
			&row.VolumesTotal,
		); err != nil {
			return nil, err
		}
		items = append(items, row)
	}
	return items, rows.Err()
}
