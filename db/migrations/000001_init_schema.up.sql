CREATE TABLE "users" (
  "user_id" uuid PRIMARY KEY,
  "nickname" VARCHAR(250),
  "first_name" VARCHAR(250),
  "last_name" VARCHAR(250),
  "email" VARCHAR(250),
  "password" VARCHAR(250),
  "is_public" boolean,
  "user_parent" uuid,
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp,
  "login_date" timestamp
);