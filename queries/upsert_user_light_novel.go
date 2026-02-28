package queries

import (
	"context"
	_ "embed"
)

//go:embed upsert_user_light_novel.sql
var upsertUserLightNovelSQL string

type UpsertUserLightNovelParameters struct {
	AnilistMediaID  int64
	ChaptersCurrent int64
	Favorite        int64
	VolumesCurrent  int64
}

func (q *Queries) UpsertUserLightNovel(c context.Context, p UpsertUserLightNovelParameters) error {
	_, err := q.connection.ExecContext(c, upsertUserLightNovelSQL,
		p.AnilistMediaID,
		p.ChaptersCurrent,
		p.Favorite,
		p.VolumesCurrent,
	)
	return err
}
