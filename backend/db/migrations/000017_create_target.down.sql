ALTER TABLE
    product DROP FOREIGN KEY product_ibfk_2;

DROP INDEX target_id ON product;

ALTER TABLE
    product DROP COLUMN target_id;

DROP TABLE IF EXISTS target;