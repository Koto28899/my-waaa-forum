CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS "sessions" (
  "id" uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4 (),
  "role_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()',
  "expires_at" timestamptz NOT NULL DEFAULT 'now()',
  "user_agent" varchar NOT NULL,
  "client_ip" varchar NOT NULL,
  "is_blocked" boolean NOT NULL DEFAULT false
);

ALTER TABLE "sessions" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("id");