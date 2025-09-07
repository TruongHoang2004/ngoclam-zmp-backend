-- Modify "image_related" table
ALTER TABLE "public"."image_related" DROP CONSTRAINT "fk_categories_image_related", DROP CONSTRAINT "fk_products_images";
