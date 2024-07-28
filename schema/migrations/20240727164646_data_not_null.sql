-- Modify "message" table
ALTER TABLE "public"."message" ALTER COLUMN "created_at" SET NOT NULL, ALTER COLUMN "updated_at" SET NOT NULL;
