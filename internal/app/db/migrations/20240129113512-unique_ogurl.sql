
-- +migrate Up
CREATE UNIQUE INDEX IF NOT EXISTS url_unique_ogurl on url (og_url);
-- +migrate Down
DROP INDEX IF EXISTS url_unique_ogurl;