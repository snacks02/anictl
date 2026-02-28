package queries

import (
	"context"
	_ "embed"
)

//go:embed upsert_user_anime.sql
var upsertUserAnimeSQL string

type UpsertUserAnimeParameters struct {
	AnilistMediaID  int64
	Favorite        int64
	ProgressCurrent int64
	Repeat          int64
}

func (q *Queries) UpsertUserAnime(c context.Context, p UpsertUserAnimeParameters) error {
	_, err := q.connection.ExecContext(c, upsertUserAnimeSQL,
		p.AnilistMediaID,
		p.Favorite,
		p.ProgressCurrent,
		p.Repeat,
	)
	return err
}
