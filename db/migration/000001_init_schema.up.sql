CREATE TABLE "users" (
  "id" bigint PRIMARY KEY,
  "username" varchar NOT NULL,
  "hashed_password" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "avatar" varchar NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT('0001-01-01 00:00:00Z'),  
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "subs" (
  "id" bigint PRIMARY KEY,
  "creator_id" bigint NOT NULL,
  "name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "mods" (
  "id" bigint PRIMARY KEY,
  "sub_id" bigint NOT NULL,
  "user_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "subscribers" (
  "id" bigint PRIMARY KEY,
  "sub_id" bigint NOT NULL,
  "user_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "posts" (
  "id" bigserial PRIMARY KEY,
  "poster_id" bigint NOT NULL,
  "sub_id" bigint NOT NULL,
  "title" varchar NOT NULL,
  "description" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "comments" (
  "id" bigserial PRIMARY KEY,
  "commenter_id" bigint NOT NULL,
  "parent_comment_id" bigint NULL,
  "post_id" bigint NOT NULL,
  "points" bigint NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "subs" ADD FOREIGN KEY ("creator_id") REFERENCES "users" ("id");

ALTER TABLE "mods" ADD FOREIGN KEY ("sub_id") REFERENCES "subs" ("id");

ALTER TABLE "mods" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "subscribers" ADD FOREIGN KEY ("sub_id") REFERENCES "subs" ("id");

ALTER TABLE "subscribers" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "posts" ADD FOREIGN KEY ("poster_id") REFERENCES "users" ("id");

ALTER TABLE "posts" ADD FOREIGN KEY ("sub_id") REFERENCES "subs" ("id");

ALTER TABLE "comments" ADD FOREIGN KEY ("commenter_id") REFERENCES "users" ("id");

ALTER TABLE "comments" ADD FOREIGN KEY ("post_id") REFERENCES "posts" ("id");

ALTER TABLE "comments" ADD FOREIGN KEY ("parent_comment_id") REFERENCES "comments" ("id");


-- CREATE INDEX ON "accounts" ("owner");

-- CREATE INDEX ON "entries" ("account_id");

-- CREATE INDEX ON "transfers" ("from_account_id");

-- CREATE INDEX ON "transfers" ("to_account_id");

-- CREATE INDEX ON "transfers" ("from_account_id", "to_account_id");

-- COMMENT ON COLUMN "entries"."amount" IS 'can be negative or positive';

COMMENT ON COLUMN "comments"."parent_comment_id" IS 'may be null for comments that are top level';
