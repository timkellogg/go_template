-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE users (
  id serial primary key,
  email text not null unique
);

-- +migrate Down
DROP TABLE users;

