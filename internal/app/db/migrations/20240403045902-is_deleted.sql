
-- +migrate Up
ALTER TABLE IF EXISTS url ADD IF NOT EXISTS is_deleted BOOLEAN;
-- +migrate Down
ALTER TABLE IF EXISTS url DROP IF EXISTS is_deleted;
