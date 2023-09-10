
CREATE TABLE "role_permission" (
  "role_id" bigserial,
  "permission_id" bigserial
);

ALTER TABLE "role_permission" ADD FOREIGN KEY ("permission_id") REFERENCES "permission" ("id");
ALTER TABLE "role_permission" ADD FOREIGN KEY ("role_id") REFERENCES "role" ("id");