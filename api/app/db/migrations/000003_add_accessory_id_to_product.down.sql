ALTER TABLE
    product DROP FOREIGN KEY product_ibfk_1;

ALTER TABLE
    product DROP COLUMN category_id;

DROP TABLE IF EXISTS category;