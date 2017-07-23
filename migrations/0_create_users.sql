-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE users (
  id serial primary key,
  email text not null unique,
  created_at timestamp not null,
  updated_at timestamp,
  deleted_at timestamp
);

-- +migrate Down
DROP TABLE users;
