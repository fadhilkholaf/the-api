CREATE TYPE ROLE AS ENUM ('user', 'admin');

CREATE TABLE IF NOT EXISTS "users" (
    "id" SERIAL PRIMARY KEY UNIQUE NOT NULL,
    "username" TEXT UNIQUE NOT NULL,
    "password" TEXT NOT NULL,
    "role" ROLE DEFAULT 'user' NOT NULL,
    "createdAt" TIMESTAMPTZ DEFAULT NOW() NOT NULL,
    "updatedAt" TIMESTAMPTZ,
    "deletedAt" TIMESTAMPTZ
);

CREATE UNIQUE INDEX IF NOT EXISTS "user_id_index" ON "users"("id");

CREATE UNIQUE INDEX IF NOT EXISTS "user_username_index" ON "users"("username");

CREATE INDEX IF NOT EXISTS "user_deletedAt_index" ON "users"("deletedAt");