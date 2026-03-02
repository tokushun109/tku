ALTER TABLE target
MODIFY COLUMN created_at DATETIME NULL,
MODIFY COLUMN updated_at DATETIME NULL;

ALTER TABLE contact
MODIFY COLUMN created_at DATETIME NULL,
MODIFY COLUMN updated_at DATETIME NULL;

ALTER TABLE session
MODIFY COLUMN created_at DATETIME NULL;

ALTER TABLE user
MODIFY COLUMN created_at DATETIME NULL,
MODIFY COLUMN updated_at DATETIME NULL;

ALTER TABLE sns
MODIFY COLUMN created_at DATETIME NULL,
MODIFY COLUMN updated_at DATETIME NULL;

ALTER TABLE skill_market
MODIFY COLUMN created_at DATETIME NULL,
MODIFY COLUMN updated_at DATETIME NULL;

ALTER TABLE creator
MODIFY COLUMN created_at DATETIME NULL,
MODIFY COLUMN updated_at DATETIME NULL;

ALTER TABLE site_detail
MODIFY COLUMN created_at DATETIME NULL,
MODIFY COLUMN updated_at DATETIME NULL;

ALTER TABLE sales_site
MODIFY COLUMN created_at DATETIME NULL,
MODIFY COLUMN updated_at DATETIME NULL;

ALTER TABLE product_image
MODIFY COLUMN created_at DATETIME NULL,
MODIFY COLUMN updated_at DATETIME NULL;

ALTER TABLE product_to_tag
MODIFY COLUMN created_at DATETIME NULL,
MODIFY COLUMN updated_at DATETIME NULL;

ALTER TABLE tag
MODIFY COLUMN created_at DATETIME NULL,
MODIFY COLUMN updated_at DATETIME NULL;

ALTER TABLE category
MODIFY COLUMN created_at DATETIME NULL,
MODIFY COLUMN updated_at DATETIME NULL;

ALTER TABLE product
MODIFY COLUMN created_at DATETIME NULL,
MODIFY COLUMN updated_at DATETIME NULL;

DROP INDEX idx_contact_deleted_at_created_at ON contact;

DROP INDEX idx_product_image_product_uuid_deleted_at_display_order_id ON product_image;

DROP INDEX idx_product_deleted_at_is_active_is_recommend_created_at_id ON product;

DROP INDEX idx_product_deleted_at_created_at_id ON product;
