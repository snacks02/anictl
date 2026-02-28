package requests

const (
	AnilistEndpoint = "https://graphql.anilist.co"
	AnilistUsername = "snacks64"
)

type MediaFormat string
type MediaListStatus string
type MediaType string

const (
	MediaFormatNovel MediaFormat = "NOVEL"
	MediaListStatusCompleted MediaListStatus = "COMPLETED"
	MediaListStatusRepeating MediaListStatus = "REPEATING"
	MediaTypeAnime MediaType = "ANIME"
	MediaTypeManga MediaType = "MANGA"
)
