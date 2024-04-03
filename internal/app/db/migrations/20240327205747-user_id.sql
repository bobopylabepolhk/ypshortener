
-- +migrate Up
ALTER TABLE IF EXISTS url ADD IF NOT EXISTS user_id VARCHAR;
-- +migrate Down
ALTER TABLE IF EXISTS url DROP IF NOT EXISTS user_id;
