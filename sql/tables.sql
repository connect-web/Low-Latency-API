CREATE SCHEMA IF NOT EXISTS users;

CREATE TABLE IF NOT EXISTS users.accs
(
    ID       SERIAL PRIMARY KEY,
    name     CITEXT UNIQUE NOT NULL,
    password BYTEA         NOT NULL
);