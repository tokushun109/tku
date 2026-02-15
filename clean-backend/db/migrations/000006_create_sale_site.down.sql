ALTER TABLE
    site_detail DROP FOREIGN KEY site_detail_ibfk_1;

ALTER TABLE
    site_detail DROP FOREIGN KEY site_detail_ibfk_2;

DROP TABLE IF EXISTS site_detail;

DROP TABLE IF EXISTS sales_site;