ALTER TABLE
    product DROP FOREIGN KEY product_ibfk_2;

ALTER TABLE
    product DROP COLUMN product_image_id;

DROP TABLE IF EXISTS product_image;