ALTER TABLE session
DROP FOREIGN KEY fk_session_user_uuid;

DROP INDEX idx_session_user_uuid ON session;

ALTER TABLE session
DROP COLUMN user_uuid,
MODIFY COLUMN user_id INT NOT NULL;

ALTER TABLE site_detail
DROP FOREIGN KEY fk_site_detail_product_uuid;

ALTER TABLE site_detail
DROP FOREIGN KEY fk_site_detail_sales_site_uuid;

DROP INDEX idx_site_detail_product_uuid ON site_detail;
DROP INDEX idx_site_detail_sales_site_uuid ON site_detail;

ALTER TABLE site_detail
DROP COLUMN product_uuid,
DROP COLUMN sales_site_uuid,
MODIFY COLUMN product_id INT NOT NULL,
MODIFY COLUMN sales_site_id INT NOT NULL;

ALTER TABLE product_image
DROP FOREIGN KEY fk_product_image_product_uuid;

DROP INDEX idx_product_image_product_uuid ON product_image;

ALTER TABLE product_image
DROP COLUMN product_uuid;

ALTER TABLE product_to_tag
DROP FOREIGN KEY fk_product_to_tag_product_uuid;

ALTER TABLE product_to_tag
DROP FOREIGN KEY fk_product_to_tag_tag_uuid;

DROP INDEX idx_product_to_tag_product_uuid ON product_to_tag;
DROP INDEX idx_product_to_tag_tag_uuid ON product_to_tag;
DROP INDEX idx_product_to_tag_product_uuid_tag_uuid ON product_to_tag;

ALTER TABLE product_to_tag
DROP COLUMN product_uuid,
DROP COLUMN tag_uuid,
MODIFY COLUMN product_id INT NOT NULL,
MODIFY COLUMN tag_id INT NOT NULL;

ALTER TABLE product
DROP FOREIGN KEY fk_product_category_uuid;

ALTER TABLE product
DROP FOREIGN KEY fk_product_target_uuid;

DROP INDEX idx_product_category_uuid ON product;
DROP INDEX idx_product_target_uuid ON product;

ALTER TABLE product
DROP COLUMN category_uuid,
DROP COLUMN target_uuid;

DROP INDEX idx_contact_uuid ON contact;

ALTER TABLE contact
DROP COLUMN uuid;

DROP INDEX idx_creator_uuid ON creator;

ALTER TABLE creator
DROP COLUMN uuid;

ALTER TABLE user
MODIFY COLUMN uuid VARCHAR(36) NOT NULL;

ALTER TABLE sales_site
MODIFY COLUMN uuid VARCHAR(36) NOT NULL;

ALTER TABLE tag
MODIFY COLUMN uuid VARCHAR(36) NOT NULL;

ALTER TABLE product
MODIFY COLUMN uuid VARCHAR(36) NOT NULL;

ALTER TABLE target
MODIFY COLUMN uuid VARCHAR(36) NOT NULL;

ALTER TABLE category
MODIFY COLUMN uuid VARCHAR(36) NOT NULL;
