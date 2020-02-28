create table posts
(
    id bigserial not null
        constraint table_name_pk
            primary key,
    user_id bigint
        constraint user_id
            references users
            on update cascade on delete cascade,
    text text,
    title varchar
);

