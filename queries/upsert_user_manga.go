package queries

import (
	"context"
	_ "embed"
)

//go:embed upsert_user_manga.sql
var upsertUserMangaSQL string

type UpsertUserMangaParameters struct {
	AnilistMediaID  int64
	ChaptersCurrent int64
	Favorite        int64
	VolumesCurrent  int64
}

func (q *Queries) UpsertUserManga(c context.Context, p UpsertUserMangaParameters) error {
	_, err := q.connection.ExecContext(c, upsertUserMangaSQL,
		p.AnilistMediaID,
		p.ChaptersCurrent,
		p.Favorite,
		p.VolumesCurrent,
	)
	return err
}
