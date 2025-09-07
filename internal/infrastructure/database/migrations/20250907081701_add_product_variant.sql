-- Modify "products" table
ALTER TABLE "public"."products" ALTER COLUMN "price" TYPE bigint;
-- Create "variant_models" table
CREATE TABLE "public"."variant_models" (
  "id" bigserial NOT NULL,
  "product_id" bigint NOT NULL,
  "sku" character varying(100) NOT NULL,
  "price" bigint NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "uni_variant_models_sku" UNIQUE ("sku"),
  CONSTRAINT "fk_products_variants" FOREIGN KEY ("product_id") REFERENCES "public"."products" ("id") ON UPDATE CASCADE ON DELETE CASCADE
);
-- Create index "idx_variant_models_deleted_at" to table: "variant_models"
CREATE INDEX "idx_variant_models_deleted_at" ON "public"."variant_models" ("deleted_at");
-- Create index "idx_variant_models_product_id" to table: "variant_models"
CREATE INDEX "idx_variant_models_product_id" ON "public"."variant_models" ("product_id");
