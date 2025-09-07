-- Create "categories" table
CREATE TABLE "public"."categories" (
  "id" bigserial NOT NULL,
  "name" character varying(100) NOT NULL,
  "description" text NULL,
  "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamptz NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_categories_deleted_at" to table: "categories"
CREATE INDEX "idx_categories_deleted_at" ON "public"."categories" ("deleted_at");
-- Create "images" table
CREATE TABLE "public"."images" (
  "id" bigserial NOT NULL,
  "path" character varying(255) NOT NULL,
  "hash" character varying(64) NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_images_hash" to table: "images"
CREATE UNIQUE INDEX "idx_images_hash" ON "public"."images" ("hash");
-- Create "products" table
CREATE TABLE "public"."products" (
  "id" bigserial NOT NULL,
  "name" character varying(255) NOT NULL,
  "description" text NULL,
  "price" numeric NOT NULL,
  "category_id" bigint NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_categories_products" FOREIGN KEY ("category_id") REFERENCES "public"."categories" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_products_category_id" to table: "products"
CREATE INDEX "idx_products_category_id" ON "public"."products" ("category_id");
-- Create "image_related" table
CREATE TABLE "public"."image_related" (
  "id" bigserial NOT NULL,
  "image_id" bigint NOT NULL,
  "entity_id" bigint NOT NULL,
  "entity_type" text NOT NULL,
  "order" bigint NULL DEFAULT 0,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_categories_image_related" FOREIGN KEY ("entity_id") REFERENCES "public"."categories" ("id") ON UPDATE CASCADE ON DELETE CASCADE,
  CONSTRAINT "fk_products_images" FOREIGN KEY ("entity_id") REFERENCES "public"."products" ("id") ON UPDATE CASCADE ON DELETE CASCADE
);
