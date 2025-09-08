-- Modify "images" table
ALTER TABLE "public"."images" DROP COLUMN "path", ALTER COLUMN "hash" SET NOT NULL, ADD COLUMN "url" character varying(512) NOT NULL, ADD COLUMN "ik_file_id" character varying(128) NOT NULL;
-- Create index "idx_images_ik_file_id" to table: "images"
CREATE UNIQUE INDEX "idx_images_ik_file_id" ON "public"."images" ("ik_file_id");
