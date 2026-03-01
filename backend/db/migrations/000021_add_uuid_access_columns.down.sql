UPDATE session s
INNER JOIN user u ON u.uuid = s.user_uuid
SET s.user_id = u.id
WHERE s.user_id IS NULL;

ALTER TABLE session
DROP FOREIGN KEY fk_session_user_uuid;

DROP INDEX idx_session_user_uuid ON session;

ALTER TABLE session
MODIFY COLUMN user_id INT NOT NULL,
DROP COLUMN user_uuid;

ALTER TABLE site_detail
DROP FOREIGN KEY fk_site_detail_product_uuid;

ALTER TABLE site_detail
DROP FOREIGN KEY fk_site_detail_sales_site_uuid;

UPDATE site_detail sd
INNER JOIN product p ON p.uuid = sd.product_uuid
SET sd.product_id = p.id
WHERE sd.product_id IS NULL;

UPDATE site_detail sd
INNER JOIN sales_site ss ON ss.uuid = sd.sales_site_uuid
SET sd.sales_site_id = ss.id
WHERE sd.sales_site_id IS NULL;

DROP INDEX idx_site_detail_product_uuid ON site_detail;
DROP INDEX idx_site_detail_sales_site_uuid ON site_detail;

ALTER TABLE site_detail
MODIFY COLUMN product_id INT NOT NULL,
MODIFY COLUMN sales_site_id INT NOT NULL,
DROP COLUMN product_uuid,
DROP COLUMN sales_site_uuid;

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

UPDATE product_to_tag ptt
INNER JOIN product p ON p.uuid = ptt.product_uuid
SET ptt.product_id = p.id
WHERE ptt.product_id IS NULL;

UPDATE product_to_tag ptt
INNER JOIN tag t ON t.uuid = ptt.tag_uuid
SET ptt.tag_id = t.id
WHERE ptt.tag_id IS NULL;

ALTER TABLE product_to_tag
MODIFY COLUMN product_id INT NOT NULL,
MODIFY COLUMN tag_id INT NOT NULL,
DROP COLUMN product_uuid,
DROP COLUMN tag_uuid;

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
