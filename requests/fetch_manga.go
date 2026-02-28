package requests

import (
	"context"

	"github.com/hasura/go-graphql-client"
)

type FetchMangaEntry struct {
	MediaID         int
	Title           string
	Format          MediaFormat
	Chapters        int
	Volumes         int
	Progress        int
	ProgressVolumes int
}

func FetchManga(c context.Context, client *graphql.Client) ([]FetchMangaEntry, error) {
	var query struct {
		MediaListCollection struct {
			Lists []struct {
				Entries []struct {
					Media struct {
						Chapters int
						Format   MediaFormat
						Id       int
						Title    struct {
							UserPreferred string
						}
						Volumes int
					}
					Progress        int
					ProgressVolumes int
				}
			}
		} `graphql:"MediaListCollection(status_in: $statusIn, type: $mediaType, userName: $userName)"`
	}
	variables := map[string]any{
		"mediaType": MediaTypeManga,
		"statusIn":  []MediaListStatus{MediaListStatusCompleted, MediaListStatusRepeating},
		"userName":  graphql.String(AnilistUsername),
	}
	if err := client.Query(c, &query, variables); err != nil {
		return nil, err
	}
	var entries []FetchMangaEntry
	for _, list := range query.MediaListCollection.Lists {
		for _, entry := range list.Entries {
			entries = append(entries, FetchMangaEntry{
				Chapters:        entry.Media.Chapters,
				Format:          entry.Media.Format,
				MediaID:         entry.Media.Id,
				Progress:        entry.Progress,
				ProgressVolumes: entry.ProgressVolumes,
				Title:           entry.Media.Title.UserPreferred,
				Volumes:         entry.Media.Volumes,
			})
		}
	}
	return entries, nil
}
