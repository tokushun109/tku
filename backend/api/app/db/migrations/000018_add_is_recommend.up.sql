ALTER TABLE
    product
ADD
    is_recommend BOOLEAN NOT NULL DEFAULT 0
AFTER
    price;