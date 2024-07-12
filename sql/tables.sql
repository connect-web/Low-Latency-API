CREATE SCHEMA IF NOT EXISTS users;

CREATE TABLE IF NOT EXISTS users.accs
(
    ID       SERIAL PRIMARY KEY,
    name     CITEXT UNIQUE NOT NULL,
    password BYTEA         NOT NULL
);


CREATE SCHEMA IF NOT EXISTS PROFILES;
-- don't create global stats yet, we don't have enough stats stored to make it worth it.

CREATE TABLE IF NOT EXISTS Profiles.global_stats
(
    Bans int
);


CREATE TABLE IF NOT EXISTS Profiles.user_stats
(
    ID INTEGER PRIMARY KEY,
    bots_tracked int,
    bots_banned int,
    bots_playerIds INT[],
    banned_experience BIGINT,
    players_added int, -- the number of users that were not in the system but the profile requested and were approved.
    FOREIGN KEY (ID) REFERENCES USERS.ACCS (id)
);

CREATE TABLE IF NOT EXISTS Profiles.user_player_requests
(
    ID INTEGER PRIMARY KEY,
    REQUESTED_ACCOUNTS text[],
    TIME timestamptz,
    FOREIGN KEY (ID) REFERENCES USERS.ACCS (id)
);