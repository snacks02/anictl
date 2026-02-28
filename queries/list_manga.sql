select
    user_manga.chapters_current,
    anilist_media.chapters as chapters_total,
    user_manga.favorite,
    anilist_media.title,
    user_manga.volumes_current,
    anilist_media.volumes as volumes_total
from user_manga
inner join anilist_media on user_manga.anilist_media_id = anilist_media.id
order by user_manga.favorite desc, anilist_media.title asc
