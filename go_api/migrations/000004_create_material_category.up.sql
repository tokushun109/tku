CREATE TABLE IF NOT EXISTS material_category(
    id INT PRIMARY KEY AUTO_INCREMENT,
    NAME VARCHAR(30),
    created_at DATETIME,
    updated_at DATETIME
);

CREATE TABLE IF NOT EXISTS product_to_material_category(
    id INT PRIMARY KEY AUTO_INCREMENT,
    product_id INT,
    material_category_id INT,
    created_at DATETIME,
    FOREIGN KEY (product_id) REFERENCES product (id),
    FOREIGN KEY (material_category_id) REFERENCES material_category (id)
);