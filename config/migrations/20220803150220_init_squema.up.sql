CREATE TABLE IF NOT EXISTS "rooms" (
    "id" bigserial PRIMARY KEY,
    "name" varchar(50) NOT NULL,
    "created_at" date NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS "chatters" (
    "username" varchar PRIMARY KEY,
    "password" varchar NOT NULL,
    "email"varchar(70) UNIQUE NOT NULL,
    "created_at" date NOT NULL DEFAULT (now())
);