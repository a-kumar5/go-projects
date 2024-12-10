CREATE TABLE "users" (
  "id" BIGSERIAL PRIMARY KEY,
  "email" varchar(256) UNIQUE NOT NULL,
  "hased_password" varchar(256) NOT NULL,
  "active" varchar(256) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);