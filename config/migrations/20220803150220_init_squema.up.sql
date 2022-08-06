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

CREATE TABLE IF NOT EXISTS "chats" (
    "id" bigserial PRIMARY KEY,
    "room" varchar NOT NULL,
    "author" varchar NOT NULL,
    "message" varchar NOT NULL,
    "created_at" date NOT NULL DEFAULT (now())
);

ALTER TABLE "chats" ADD FOREIGN KEY ("author") REFERENCES "chatters" ("username");