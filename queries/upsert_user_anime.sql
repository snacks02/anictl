insert into user_anime (
    anilist_media_id,
    favorite,
    progress_current,
    repeat
)
values (?, ?, ?, ?)
on conflict (anilist_media_id) do update set
    favorite = excluded.favorite,
    progress_current = excluded.progress_current,
    repeat = excluded.repeat
