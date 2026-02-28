ALTER TABLE
    product
ADD
    category_id INT NULL;

ALTER TABLE
    product
ADD
    FOREIGN KEY (category_id) REFERENCES category(id);