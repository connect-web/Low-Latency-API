CREATE SCHEMA IF NOT EXISTS users;

CREATE TABLE IF NOT EXISTS users.accs
(
    ID       SERIAL PRIMARY KEY,
    TIME     timestamptz DEFAULT NOW(),
    name     CITEXT UNIQUE NOT NULL,
    password BYTEA         NOT NULL
);



CREATE TABLE IF NOT EXISTS users.fiber_storage
(
    k VARCHAR(64) PRIMARY KEY NOT NULL DEFAULT '',
    v BYTEA                   NOT NULL,
    e BIGINT                  NOT NULL DEFAULT '0'
);

CREATE SCHEMA IF NOT EXISTS PROFILES;
-- don't create global stats yet, we don't have enough stats stored to make it worth it.

CREATE TABLE IF NOT EXISTS Profiles.global_stats
(
    Bans int
);


CREATE TABLE IF NOT EXISTS Profiles.user_stats
(
    ID                INTEGER PRIMARY KEY,
    TIME              timestamptz DEFAULT NOW(),
    bots_tracked      int         DEFAULT 0,
    bots_banned       int         DEFAULT 0,
    bots_playerIds    INT[]       DEFAULT ARRAY []::integer[],
    banned_experience BIGINT      default 0,
    players_added     int         default 0, -- the number of users that were not in the system but the profile requested and were approved.
    FOREIGN KEY (ID) REFERENCES USERS.ACCS (id)
);

CREATE TABLE IF NOT EXISTS Profiles.user_player_requests
(
    ID                 INTEGER PRIMARY KEY,
    REQUESTED_ACCOUNTS text[]      DEFAULT ARRAY []::TEXT[],
    TIME               timestamptz DEFAULT NOW(),
    FOREIGN KEY (ID) REFERENCES USERS.ACCS (id)
);


