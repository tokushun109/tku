CREATE TABLE IF NOT EXISTS sale_site(
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(30),
    url VARCHAR(255),
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME
);

CREATE TABLE IF NOT EXISTS product_to_sale_site(
    id INT PRIMARY KEY AUTO_INCREMENT,
    product_id INT,
    sale_site_id INT,
    created_at DATETIME,
    deleted_at DATETIME,
    FOREIGN KEY (product_id) REFERENCES product (id),
    FOREIGN KEY (sale_site_id) REFERENCES sale_site(id)
);