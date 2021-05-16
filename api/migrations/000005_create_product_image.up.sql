CREATE TABLE IF NOT EXISTS product_image(
    id INT PRIMARY KEY AUTO_INCREMENT,
    uuid VARCHAR(36) NOT NULL UNIQUE,
    name VARCHAR(30),
    mime_type VARCHAR(30),
    path VARCHAR(255),
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME
);

ALTER TABLE
    product_image
ADD
    product_id INT NULL;

ALTER TABLE
    product_image
ADD
    FOREIGN KEY (product_id) REFERENCES product(id);