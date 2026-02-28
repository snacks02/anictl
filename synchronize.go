package main

import (
	"anictl/queries"
	"anictl/requests"
	"context"
	"database/sql"
	"net/http"

	"github.com/hasura/go-graphql-client"
)

func synchronize(c context.Context, db *sql.DB, q *queries.Queries) error {
	client := graphql.NewClient(requests.AnilistEndpoint, http.DefaultClient)

	tx, err := db.BeginTx(c, nil)
	if err != nil {
		return err
	}
	txQueries := q.WithTx(tx)

	favorites, err := requests.FetchFavorites(c, client)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := synchronizeAnime(c, client, txQueries, favorites); err != nil {
		tx.Rollback()
		return err
	}
	if err := synchronizeManga(c, client, txQueries, favorites); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func synchronizeAnime(c context.Context, client *graphql.Client, q *queries.Queries, favorites map[int]bool) error {
	entries, err := requests.FetchAnime(c, client)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		if err := q.UpsertAnilistMedia(c, queries.UpsertAnilistMediaParameters{
			Episodes: nullInt64(entry.Episodes),
			Format:   string(entry.Format),
			ID:       int64(entry.MediaID),
			Title:    entry.Title,
		}); err != nil {
			return err
		}
		if err := q.UpsertUserAnime(c, queries.UpsertUserAnimeParameters{
			AnilistMediaID:  int64(entry.MediaID),
			Favorite:        boolToInt64(favorites[entry.MediaID]),
			ProgressCurrent: int64(entry.Progress),
			Repeat:          int64(entry.Repeat),
		}); err != nil {
			return err
		}
	}
	return nil
}

func synchronizeManga(c context.Context, client *graphql.Client, q *queries.Queries, favorites map[int]bool) error {
	entries, err := requests.FetchManga(c, client)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		if err := q.UpsertAnilistMedia(c, queries.UpsertAnilistMediaParameters{
			Chapters: nullInt64(entry.Chapters),
			Format:   string(entry.Format),
			ID:       int64(entry.MediaID),
			Title:    entry.Title,
			Volumes:  nullInt64(entry.Volumes),
		}); err != nil {
			return err
		}
		progress := queries.UpsertUserMangaParameters{
			AnilistMediaID:  int64(entry.MediaID),
			ChaptersCurrent: int64(entry.Progress),
			Favorite:        boolToInt64(favorites[entry.MediaID]),
			VolumesCurrent:  int64(entry.ProgressVolumes),
		}
		if entry.Format == requests.MediaFormatNovel {
			if err := q.UpsertUserLightNovel(c, queries.UpsertUserLightNovelParameters(progress)); err != nil {
				return err
			}
		} else {
			if err := q.UpsertUserManga(c, progress); err != nil {
				return err
			}
		}
	}
	return nil
}

func boolToInt64(b bool) int64 {
	if b {
		return 1
	}
	return 0
}

func nullInt64(n int) sql.NullInt64 {
	if n == 0 {
		return sql.NullInt64{}
	}
	return sql.NullInt64{Valid: true, Int64: int64(n)}
}
