-- +goose Up
CREATE TABLE chat (
    id serial primary key,
    name text not null,
    from_user text not null
);


-- +goose Down
DROP TABLE chat;

