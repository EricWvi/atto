use atto_dev;

drop table if exists songs;
create table songs
(
    id         int auto_increment
        primary key,
    sid        bigint default 0 not null,
    album_id     bigint default 0 not null,
    artist_id     bigint default 0 not null,
    title       varchar(200)     null,
    comment       varchar(400)     null,
    duration     float(8, 2)     not null,
    track      int          null,
    disc       int     default 1  not null,
    link       varchar(240) default '' not null,
    cover       varchar(240) default '' not null,
    created_at timestamp        null,
    updated_at timestamp        null,
    deleted_at timestamp        null
)
    collate = utf8mb4_unicode_ci;

create index album_id__index
    on songs (album_id);

create index artist_id__index
    on songs (artist_id);

create index sid__index
    on songs (sid);



