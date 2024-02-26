-- +goose Up
create table users(
    id serial primary key,
    created_at timestamptz not null,
    updated_at timestamptz not null,
    first_name text not null unique,
    last_name text
);

-- +goose Down
drop table users;