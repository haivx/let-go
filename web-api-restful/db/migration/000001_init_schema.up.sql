
CREATE TABLE "role" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "name" varchar DEFAULT 'USER'
);


CREATE TABLE "user" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "username" varchar NOT NULL,
  "email" varchar NOT NULL,
  "phone_number" varchar,
  "password" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz
);


CREATE TABLE "user_role" (
  "user_id" bigserial NOT NULL,
  "role_id" bigserial NOT NULL
);


ALTER TABLE "user_role" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");
ALTER TABLE "user_role" ADD FOREIGN KEY ("role_id") REFERENCES "role" ("id");


-- CREATE INDEX ON "users" ("username");