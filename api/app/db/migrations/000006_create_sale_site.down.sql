ALTER TABLE
    product_to_sales_site DROP FOREIGN KEY product_to_sales_site_ibfk_1;

ALTER TABLE
    product_to_sales_site DROP FOREIGN KEY product_to_sales_site_ibfk_2;

DROP TABLE IF EXISTS product_to_sales_site;

DROP TABLE IF EXISTS sales_site;