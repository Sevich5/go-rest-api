-- Create "users" table
CREATE TABLE "users" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "email" text NULL,
  "password" text NULL,
  "created_at" timestamptz NULL,
  PRIMARY KEY ("id")
);
