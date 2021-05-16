ALTER TABLE
    product_image DROP FOREIGN KEY product_image_ibfk_1;

ALTER TABLE
    product_image DROP COLUMN product_id;

DROP TABLE IF EXISTS product_image;