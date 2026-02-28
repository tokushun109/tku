CREATE TABLE IF NOT EXISTS target(
    id INT PRIMARY KEY AUTO_INCREMENT,
    uuid VARCHAR(36) NOT NULL UNIQUE,
    name VARCHAR(30) NOT NULL,
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME
);

ALTER TABLE
    product
ADD
    target_id INT NULL
AFTER
    category_id;

ALTER TABLE
    product
ADD
    FOREIGN KEY (target_id) REFERENCES target(id);