-- +goose Up
create table products(
    id serial primary key,
    created_at timestamptz not null,
    updated_at timestamptz not null,
    name text not null unique,
    serial_number text
);

-- +goose Down
drop table products;