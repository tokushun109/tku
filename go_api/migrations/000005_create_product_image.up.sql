CREATE TABLE IF NOT EXISTS product_image(
    id INT PRIMARY KEY AUTO_INCREMENT,
    uuid VARCHAR(36) NOT NULL UNIQUE,
    name VARCHAR(30),
    mime_type VARCHAR(30),
    path VARCHAR(255),
    created_at DATETIME,
    updated_at DATETIME
);

ALTER TABLE
    product
ADD
    product_image_id INT NULL;

ALTER TABLE
    product
ADD
    FOREIGN KEY (product_image_id) REFERENCES product_image(id);