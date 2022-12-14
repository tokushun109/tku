ALTER TABLE
    product_image
ADD
    `order` INT NOT NULL DEFAULT 0
AFTER
    path;