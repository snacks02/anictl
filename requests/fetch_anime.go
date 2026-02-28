package requests

import (
	"context"

	"github.com/hasura/go-graphql-client"
)

type FetchAnimeEntry struct {
	Episodes int
	Format   MediaFormat
	MediaID  int
	Progress int
	Repeat   int
	Title    string
}

func FetchAnime(c context.Context, client *graphql.Client) ([]FetchAnimeEntry, error) {
	var query struct {
		MediaListCollection struct {
			Lists []struct {
				Entries []struct {
					Media struct {
						Episodes int
						Format   MediaFormat
						Id    int
						Title struct {
							UserPreferred string
						}
					}
					Progress int
					Repeat   int
				}
			}
		} `graphql:"MediaListCollection(status_in: $statusIn, type: $mediaType, userName: $userName)"`
	}
	variables := map[string]any{
		"mediaType": MediaTypeAnime,
		"statusIn":  []MediaListStatus{MediaListStatusCompleted, MediaListStatusRepeating},
		"userName":  graphql.String(AnilistUsername),
	}
	if err := client.Query(c, &query, variables); err != nil {
		return nil, err
	}
	var entries []FetchAnimeEntry
	for _, list := range query.MediaListCollection.Lists {
		for _, entry := range list.Entries {
			entries = append(entries, FetchAnimeEntry{
				Episodes: entry.Media.Episodes,
				Format:   entry.Media.Format,
				MediaID:  entry.Media.Id,
				Progress: entry.Progress,
				Repeat:   entry.Repeat,
				Title:    entry.Media.Title.UserPreferred,
			})
		}
	}
	return entries, nil
}
