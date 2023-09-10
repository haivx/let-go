CREATE TABLE "permission" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "code" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);