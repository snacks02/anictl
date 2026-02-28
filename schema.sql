create table if not exists anilist_media (
    chapters integer,
    episodes integer,
    format text not null,
    id integer primary key,
    title text not null,
    volumes integer
);

create table if not exists user_anime (
    anilist_media_id integer not null unique references anilist_media (id),
    favorite integer not null default 0,
    id integer primary key autoincrement,
    progress_current integer not null default 0,
    repeat integer not null default 0
);

create table if not exists user_manga (
    anilist_media_id integer not null unique references anilist_media (id),
    chapters_current integer not null default 0,
    favorite integer not null default 0,
    id integer primary key autoincrement,
    volumes_current integer not null default 0
);

create table if not exists user_light_novels (
    anilist_media_id integer not null unique references anilist_media (id),
    chapters_current integer not null default 0,
    favorite integer not null default 0,
    id integer primary key autoincrement,
    volumes_current integer not null default 0
);
