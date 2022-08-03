CREATE TABLE IF NOT EXISTS "rooms" (
    "id" bigserial PRIMARY KEY,
    "name" varchar(50) NOT NULL,
    "created_at" date NOT NULL DEFAULT (now())
);