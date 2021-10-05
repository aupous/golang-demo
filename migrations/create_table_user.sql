-- +migrate Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT,
    email TEXT UNIQUE NOT NULL,
    password TEXT,
    full_name TEXT
);

-- +migrate Down
DROP TABLE IF EXISTS users;
DROP EXTENSION IF EXISTS "uuid-ossp";