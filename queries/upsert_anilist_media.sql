insert into anilist_media (
    chapters,
    episodes,
    format,
    id,
    title,
    volumes
)
values (?, ?, ?, ?, ?, ?)
on conflict (id) do update set
    chapters = excluded.chapters,
    episodes = excluded.episodes,
    format = excluded.format,
    title = excluded.title,
    volumes = excluded.volumes
