-- 000021 までで uuid ベースの関連が主系になった前提で、
-- product 系に残っていた旧 id 外部キーを段階的に取り除く。
-- 旧 FK は過去 migration で無名作成されており、現行スキーマ上では
-- *_ibfk_* の制約名で作られているため、その実名を指定して削除する。

-- product の category_id / target_id を削除する。
ALTER TABLE product
DROP FOREIGN KEY product_ibfk_1;

ALTER TABLE product
DROP FOREIGN KEY product_ibfk_2;

ALTER TABLE product
DROP COLUMN category_id,
DROP COLUMN target_id;

-- product_to_tag の product_id / tag_id を削除する。
ALTER TABLE product_to_tag
DROP FOREIGN KEY product_to_tag_ibfk_1;

ALTER TABLE product_to_tag
DROP FOREIGN KEY product_to_tag_ibfk_2;

ALTER TABLE product_to_tag
DROP COLUMN product_id,
DROP COLUMN tag_id;

-- product_image の product_id を削除する。
ALTER TABLE product_image
DROP FOREIGN KEY product_image_ibfk_1;

ALTER TABLE product_image
DROP COLUMN product_id;

-- site_detail の product_id / sales_site_id を削除する。
ALTER TABLE site_detail
DROP FOREIGN KEY site_detail_ibfk_1;

ALTER TABLE site_detail
DROP FOREIGN KEY site_detail_ibfk_2;

ALTER TABLE site_detail
DROP COLUMN product_id,
DROP COLUMN sales_site_id;
