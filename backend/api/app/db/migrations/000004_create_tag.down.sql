ALTER TABLE
    product_to_tag DROP FOREIGN KEY product_to_tag_ibfk_1;

ALTER TABLE
    product_to_tag DROP FOREIGN KEY product_to_tag_ibfk_2;

DROP TABLE IF EXISTS product_to_tag;

DROP TABLE IF EXISTS tag;