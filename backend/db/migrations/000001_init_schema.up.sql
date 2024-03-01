CREATE TYPE "role_types" AS ENUM (
  'user',
  'admin'
);

CREATE TYPE "notification_types" AS ENUM (
  'adopt',
  'reply',
  'like',
  'at',
  'from_admin'
);

CREATE TYPE "secne_types" AS ENUM (
  'post',
  'comment',
  'reply'
);

CREATE TABLE "roles" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()',
  "updated_at" timestamptz NOT NULL DEFAULT 'now()',
  "deleted_at" timestamptz,
  "email" varchar(321) UNIQUE NOT NULL,
  "hashed_pwd" varchar NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT 'now()',
  "role_name" varchar(21) NOT NULL,
  "type" role_types NOT NULL DEFAULT 'user',
  "statement" varchar(127),
  "posts_count" bigint NOT NULL DEFAULT 0,
  "comments_count" bigint NOT NULL DEFAULT 0,
  "replies_count" bigint NOT NULL DEFAULT 0,
  "likes_count" bigint NOT NULL DEFAULT 0,
  "helps_count" bigint NOT NULL DEFAULT 0,
  "fans_count" bigint NOT NULL DEFAULT 0
);

CREATE TABLE "follows" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()',
  "updated_at" timestamptz NOT NULL DEFAULT 'now()',
  "deleted_at" timestamptz,
  "from_role_id" bigint NOT NULL,
  "to_role_id" bigint NOT NULL
);

CREATE TABLE "section" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()',
  "updated_at" timestamptz NOT NULL DEFAULT 'now()',
  "deleted_at" timestamptz,
  "section_name" varchar(21) NOT NULL,
  "statement" varchar(127),
  "manager_id" bigint NOT NULL
);

CREATE TABLE "posts" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()',
  "updated_at" timestamptz NOT NULL DEFAULT 'now()',
  "deleted_at" timestamptz,
  "from_role_id" bigint NOT NULL,
  "to_section_id" bigint NOT NULL,
  "likes_count" bigint NOT NULL DEFAULT 0,
  "is_top" boolean NOT NULL DEFAULT false,
  "is_highlight" boolean NOT NULL DEFAULT false,
  "title" varchar(127) NOT NULL,
  "body" varchar(1023) NOT NULL
);

CREATE TABLE "votes" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()',
  "updated_at" timestamptz NOT NULL DEFAULT 'now()',
  "deleted_at" timestamptz,
  "from_post_id" bigint NOT NULL,
  "expired_at" timestamptz NOT NULL DEFAULT 'now()',
  "register" boolean NOT NULL DEFAULT false
);

CREATE TABLE "vote_options" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()',
  "updated_at" timestamptz NOT NULL DEFAULT 'now()',
  "deleted_at" timestamptz,
  "vote_id" bigint NOT NULL,
  "info" varchar(127) NOT NULL,
  "count" bigint NOT NULL DEFAULT 0
);

CREATE TABLE "vote_events" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()',
  "updated_at" timestamptz NOT NULL DEFAULT 'now()',
  "deleted_at" timestamptz,
  "from_role_id" bigint NOT NULL,
  "to_vote_id" bigint NOT NULL,
  "to_vote_option" bigint NOT NULL
);

CREATE TABLE "helps" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()',
  "updated_at" timestamptz NOT NULL DEFAULT 'now()',
  "deleted_at" timestamptz,
  "from_post_id" bigint NOT NULL,
  "adopt_comment_id" bigint
);

CREATE TABLE "comments" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()',
  "updated_at" timestamptz NOT NULL DEFAULT 'now()',
  "deleted_at" timestamptz,
  "from_role_id" bigint NOT NULL,
  "to_post_id" bigint NOT NULL,
  "likes_count" bigint NOT NULL DEFAULT 0,
  "is_top" boolean NOT NULL DEFAULT false,
  "body" varchar(1023) NOT NULL
);

CREATE TABLE "replies" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()',
  "updated_at" timestamptz NOT NULL DEFAULT 'now()',
  "deleted_at" timestamptz,
  "from_role_id" bigint NOT NULL,
  "to_comment_id" bigint NOT NULL,
  "likes_count" bigint NOT NULL DEFAULT 0,
  "body" varchar(1023) NOT NULL
);

CREATE TABLE "notifications" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()',
  "updated_at" timestamptz NOT NULL DEFAULT 'now()',
  "deleted_at" timestamptz,
  "notification_type" notification_types NOT NULL,
  "from_role_id" bigint NOT NULL,
  "scene_type" secne_types NOT NULL,
  "scene_id" bigint NOT NULL,
  "to_role_id" bigint NOT NULL,
  "is_checked" bool NOT NULL DEFAULT false
);

CREATE TABLE "stars" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()',
  "updated_at" timestamptz NOT NULL DEFAULT 'now()',
  "deleted_at" timestamptz,
  "from_role_id" bigint NOT NULL,
  "scene_type" secne_types NOT NULL,
  "scene_id" bigint NOT NULL
);

CREATE INDEX ON "roles" ("id");

CREATE INDEX ON "roles" ("role_name");

CREATE INDEX ON "roles" ("email");

CREATE INDEX ON "follows" ("id");

CREATE INDEX ON "follows" ("from_role_id");

CREATE INDEX ON "follows" ("to_role_id");

CREATE INDEX ON "section" ("id");

CREATE INDEX ON "posts" ("id");

CREATE INDEX ON "votes" ("from_post_id");

CREATE INDEX ON "vote_options" ("id");

CREATE INDEX ON "vote_options" ("vote_id");

CREATE INDEX ON "vote_events" ("id");

CREATE INDEX ON "vote_events" ("to_vote_id");

CREATE INDEX ON "vote_events" ("to_vote_id", "to_vote_option");

CREATE INDEX ON "helps" ("id");

CREATE INDEX ON "helps" ("from_post_id");

CREATE INDEX ON "comments" ("id");

CREATE INDEX ON "comments" ("to_post_id");

CREATE INDEX ON "replies" ("id");

CREATE INDEX ON "replies" ("to_comment_id");

CREATE INDEX ON "notifications" ("id");

CREATE INDEX ON "stars" ("id");

CREATE INDEX ON "stars" ("from_role_id");

ALTER TABLE "follows" ADD FOREIGN KEY ("from_role_id") REFERENCES "roles" ("id");

ALTER TABLE "follows" ADD FOREIGN KEY ("to_role_id") REFERENCES "roles" ("id");

ALTER TABLE "section" ADD FOREIGN KEY ("manager_id") REFERENCES "roles" ("id");

ALTER TABLE "posts" ADD FOREIGN KEY ("from_role_id") REFERENCES "roles" ("id");

ALTER TABLE "posts" ADD FOREIGN KEY ("to_section_id") REFERENCES "section" ("id");

ALTER TABLE "votes" ADD FOREIGN KEY ("from_post_id") REFERENCES "posts" ("id");

ALTER TABLE "vote_options" ADD FOREIGN KEY ("vote_id") REFERENCES "votes" ("id");

ALTER TABLE "vote_events" ADD FOREIGN KEY ("from_role_id") REFERENCES "roles" ("id");

ALTER TABLE "vote_events" ADD FOREIGN KEY ("to_vote_id") REFERENCES "votes" ("id");

ALTER TABLE "vote_events" ADD FOREIGN KEY ("to_vote_option") REFERENCES "vote_options" ("id");

ALTER TABLE "helps" ADD FOREIGN KEY ("from_post_id") REFERENCES "posts" ("id");

ALTER TABLE "helps" ADD FOREIGN KEY ("adopt_comment_id") REFERENCES "comments" ("id");

ALTER TABLE "comments" ADD FOREIGN KEY ("from_role_id") REFERENCES "roles" ("id");

ALTER TABLE "comments" ADD FOREIGN KEY ("to_post_id") REFERENCES "posts" ("id");

ALTER TABLE "replies" ADD FOREIGN KEY ("from_role_id") REFERENCES "roles" ("id");

ALTER TABLE "replies" ADD FOREIGN KEY ("to_comment_id") REFERENCES "comments" ("id");

ALTER TABLE "notifications" ADD FOREIGN KEY ("from_role_id") REFERENCES "roles" ("id");

ALTER TABLE "notifications" ADD FOREIGN KEY ("to_role_id") REFERENCES "roles" ("id");

ALTER TABLE "stars" ADD FOREIGN KEY ("from_role_id") REFERENCES "roles" ("id");
