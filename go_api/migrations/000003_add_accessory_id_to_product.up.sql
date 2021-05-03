ALTER TABLE
    product
ADD
    accessory_category_id INT NULL;

ALTER TABLE
    product
ADD
    FOREIGN KEY (accessory_category_id) REFERENCES accessory_category(id);