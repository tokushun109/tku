ALTER TABLE
    product DROP FOREIGN KEY product_ibfk_1;

ALTER TABLE
    product DROP COLUMN accessory_category_id;

DROP TABLE IF EXISTS accessory_category;