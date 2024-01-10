SELECT 'CREATE DATABASE events' WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'events')\gexec
SELECT 'CREATE DATABASE habrdb' WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'habrdb')\gexec
DO
$do$
BEGIN
    IF NOT EXISTS (SELECT FROM pg_catalog.pg_roles WHERE rolname = 'habrpguser') THEN
        CREATE ROLE habrpguser LOGIN PASSWORD 'habrpguser';
    END IF;
END
$do$;
CREATE TABLE IF NOT EXISTS events (
  "id" VARCHAR(36),
  "title" VARCHAR(256),
  "description" VARCHAR(2048),
  "user_id" INT,
  "priority" INT,
  "group" INT,
  "start" VARCHAR(25),
  "end" VARCHAR(25),
  PRIMARY KEY (id)
);
CREATE TABLE IF NOT EXISTS users (
  "id" INT,
  "name" VARCHAR(32),
  "password" VARCHAR(64),
  PRIMARY KEY (id)
);