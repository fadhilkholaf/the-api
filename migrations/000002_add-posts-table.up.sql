CREATE TABLE IF NOT EXISTS "posts"(
    "id" SERIAL PRIMARY KEY UNIQUE NOT NULL,
    "title" TEXT NOT NULL,
    "description" TEXT NOT NULL,
    "authorId" INT NOT NULL,
    "createdAt" TIMESTAMPTZ DEFAULT NOW() NOT NULL,
    "updatedAt" TIMESTAMPTZ,
    "deletedAt" TIMESTAMPTZ,
    FOREIGN KEY("authorId") REFERENCES "users"("id") ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE UNIQUE INDEX IF NOT EXISTS "post_id_index" ON "posts"("id");

CREATE INDEX IF NOT EXISTS "post_deletedAt_index" ON "posts"("deletedAt");