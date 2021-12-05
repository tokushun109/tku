CREATE TABLE IF NOT EXISTS tag(
    id INT PRIMARY KEY AUTO_INCREMENT,
    uuid VARCHAR(36) NOT NULL UNIQUE,
    name VARCHAR(30) NOT NULL,
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME
);

CREATE TABLE IF NOT EXISTS product_to_tag(
    id INT PRIMARY KEY AUTO_INCREMENT,
    product_id INT NOT NULL,
    tag_id INT NOT NULL,
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME,
    FOREIGN KEY (product_id) REFERENCES product (id),
    FOREIGN KEY (tag_id) REFERENCES tag (id)
);