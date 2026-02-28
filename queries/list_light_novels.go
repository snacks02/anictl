package queries

import (
	"context"
	"database/sql"
	_ "embed"
)

//go:embed list_light_novels.sql
var listLightNovelsSQL string

type ListLightNovelsRow struct {
	ChaptersCurrent int64
	ChaptersTotal   sql.NullInt64
	Favorite        int64
	Title           string
	VolumesCurrent  int64
	VolumesTotal    sql.NullInt64
}

func (q *Queries) ListLightNovels(c context.Context) ([]ListLightNovelsRow, error) {
	rows, err := q.connection.QueryContext(c, listLightNovelsSQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListLightNovelsRow
	for rows.Next() {
		var row ListLightNovelsRow
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
