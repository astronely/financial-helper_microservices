-- +goose Up
create table users (
    id serial primary key,
    email text not null,
    name text not null,
    password text not null,
    created_at timestamp not null default now(),
    updated_at timestamp
);

-- +goose Down
drop table user;