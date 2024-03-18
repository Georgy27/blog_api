-- +goose Up
create table post (
        id serial primary key,
        title text not null,
        short_description text not null,
        content text not null,
        blog_id int REFERENCES blog (id) not null,
        created_at timestamp not null default now(),
        updated_at timestamp
);

-- +goose Down
drop table post;