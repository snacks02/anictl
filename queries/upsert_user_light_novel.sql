insert into user_light_novels (
    anilist_media_id,
    chapters_current,
    favorite,
    volumes_current
)
values (?, ?, ?, ?)
on conflict (anilist_media_id) do update set
    chapters_current = excluded.chapters_current,
    favorite = excluded.favorite,
    volumes_current = excluded.volumes_current
