select
    user_light_novels.chapters_current,
    anilist_media.chapters as chapters_total,
    user_light_novels.favorite,
    anilist_media.title,
    user_light_novels.volumes_current,
    anilist_media.volumes as volumes_total
from user_light_novels
inner join
    anilist_media
    on user_light_novels.anilist_media_id = anilist_media.id
order by user_light_novels.favorite desc, anilist_media.title asc
