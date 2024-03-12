-- +goose Up
create table blog (
    id serial primary key,
    name text not null,
    description text not null,
    website_url text not null,
    created_at timestamp not null default now(),
    updated_at timestamp
);

-- +goose Down
drop table blog;
