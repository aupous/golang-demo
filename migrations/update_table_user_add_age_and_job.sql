-- +migrate Up
ALTER TABLE users
    ADD COLUMN age int,
    ADD COLUMN job text;

-- +migrate Down
ALTER TABLE users
    DROP COLUMN age,
    DROP COLUMN job;