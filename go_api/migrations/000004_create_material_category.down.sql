ALTER TABLE
    product_to_material_category DROP FOREIGN KEY product_to_material_category_ibfk_1;

ALTER TABLE
    product_to_material_category DROP FOREIGN KEY product_to_material_category_ibfk_2;

DROP TABLE IF EXISTS product_to_material_category;

DROP TABLE IF EXISTS material_category;