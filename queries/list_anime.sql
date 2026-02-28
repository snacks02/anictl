select
    user_anime.favorite,
    user_anime.progress_current,
    anilist_media.episodes as progress_total,
    user_anime.repeat,
    anilist_media.title
from user_anime
inner join anilist_media on user_anime.anilist_media_id = anilist_media.id
order by user_anime.favorite desc, anilist_media.title asc
