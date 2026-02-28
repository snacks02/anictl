package requests

import (
	"context"

	"github.com/hasura/go-graphql-client"
)

func FetchFavorites(c context.Context, client *graphql.Client) (map[int]bool, error) {
	var query struct {
		User struct {
			Favourites struct {
				Anime struct {
					Nodes []struct{ Id int }
				}
				Manga struct {
					Nodes []struct{ Id int }
				}
			}
		} `graphql:"User(name: $userName)"`
	}
	variables := map[string]any{
		"userName": graphql.String(AnilistUsername),
	}
	if err := client.Query(c, &query, variables); err != nil {
		return nil, err
	}
	favorites := make(map[int]bool)
	for _, node := range query.User.Favourites.Anime.Nodes {
		favorites[node.Id] = true
	}
	for _, node := range query.User.Favourites.Manga.Nodes {
		favorites[node.Id] = true
	}
	return favorites, nil
}
