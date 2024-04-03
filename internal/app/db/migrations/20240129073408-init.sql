
-- +migrate Up
CREATE TABLE url (
    id SERIAL PRIMARY KEY,
    og_url VARCHAR NOT NULL,
    short_url VARCHAR NOT NULL
);
-- +migrate Down
DROP TABLE url;
