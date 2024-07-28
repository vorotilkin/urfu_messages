-- Create "message" table
CREATE TABLE "public"."message" ("id" serial NOT NULL, "user_id" integer NOT NULL, "message" text NOT NULL, "created_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP, "updated_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP, PRIMARY KEY ("id"));
-- Create index "idx_id" to table: "message"
CREATE INDEX "idx_id" ON "public"."message" ("id");
-- Create index "idx_user_id" to table: "message"
CREATE INDEX "idx_user_id" ON "public"."message" ("user_id");
