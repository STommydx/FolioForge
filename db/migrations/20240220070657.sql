-- Create "profiles" table
CREATE TABLE "public"."profiles" (
  "id" bytea NOT NULL,
  "profile_name" text NULL,
  "first_name" text NULL,
  "last_name" text NULL,
  "email" text NULL,
  "phone" text NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_profiles_profile_name" to table: "profiles"
CREATE UNIQUE INDEX "idx_profiles_profile_name" ON "public"."profiles" ("profile_name");
