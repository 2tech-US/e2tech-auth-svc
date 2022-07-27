CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "phone" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "name" varchar NOT NULL,
  "role" varchar NOT NULL,
  "date_of_birth" date NOT NULL,
  "verified" boolean NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);