-- 中優先度の改善:
-- 1. 一覧 / カルーセル系の読み取りに効く複合 index を追加する
-- 2. created_at / updated_at を DATETIME(6) NOT NULL へ寄せる

CREATE INDEX idx_product_deleted_at_created_at_id
ON product (deleted_at, created_at, id);

CREATE INDEX idx_product_deleted_at_is_active_is_recommend_created_at_id
ON product (deleted_at, is_active, is_recommend, created_at, id);

CREATE INDEX idx_product_image_product_uuid_deleted_at_display_order_id
ON product_image (product_uuid, deleted_at, display_order, id);

CREATE INDEX idx_contact_deleted_at_created_at
ON contact (deleted_at, created_at);

-- NULL の timestamp を補完してから NOT NULL 化する。
UPDATE product
SET
    created_at = COALESCE(created_at, UTC_TIMESTAMP(6)),
    updated_at = COALESCE(updated_at, created_at, UTC_TIMESTAMP(6))
WHERE created_at IS NULL OR updated_at IS NULL;

UPDATE category
SET
    created_at = COALESCE(created_at, UTC_TIMESTAMP(6)),
    updated_at = COALESCE(updated_at, created_at, UTC_TIMESTAMP(6))
WHERE created_at IS NULL OR updated_at IS NULL;

UPDATE tag
SET
    created_at = COALESCE(created_at, UTC_TIMESTAMP(6)),
    updated_at = COALESCE(updated_at, created_at, UTC_TIMESTAMP(6))
WHERE created_at IS NULL OR updated_at IS NULL;

UPDATE product_to_tag
SET
    created_at = COALESCE(created_at, UTC_TIMESTAMP(6)),
    updated_at = COALESCE(updated_at, created_at, UTC_TIMESTAMP(6))
WHERE created_at IS NULL OR updated_at IS NULL;

UPDATE product_image
SET
    created_at = COALESCE(created_at, UTC_TIMESTAMP(6)),
    updated_at = COALESCE(updated_at, created_at, UTC_TIMESTAMP(6))
WHERE created_at IS NULL OR updated_at IS NULL;

UPDATE sales_site
SET
    created_at = COALESCE(created_at, UTC_TIMESTAMP(6)),
    updated_at = COALESCE(updated_at, created_at, UTC_TIMESTAMP(6))
WHERE created_at IS NULL OR updated_at IS NULL;

UPDATE site_detail
SET
    created_at = COALESCE(created_at, UTC_TIMESTAMP(6)),
    updated_at = COALESCE(updated_at, created_at, UTC_TIMESTAMP(6))
WHERE created_at IS NULL OR updated_at IS NULL;

UPDATE creator
SET
    created_at = COALESCE(created_at, UTC_TIMESTAMP(6)),
    updated_at = COALESCE(updated_at, created_at, UTC_TIMESTAMP(6))
WHERE created_at IS NULL OR updated_at IS NULL;

UPDATE skill_market
SET
    created_at = COALESCE(created_at, UTC_TIMESTAMP(6)),
    updated_at = COALESCE(updated_at, created_at, UTC_TIMESTAMP(6))
WHERE created_at IS NULL OR updated_at IS NULL;

UPDATE sns
SET
    created_at = COALESCE(created_at, UTC_TIMESTAMP(6)),
    updated_at = COALESCE(updated_at, created_at, UTC_TIMESTAMP(6))
WHERE created_at IS NULL OR updated_at IS NULL;

UPDATE user
SET
    created_at = COALESCE(created_at, UTC_TIMESTAMP(6)),
    updated_at = COALESCE(updated_at, created_at, UTC_TIMESTAMP(6))
WHERE created_at IS NULL OR updated_at IS NULL;

UPDATE session
SET created_at = COALESCE(created_at, UTC_TIMESTAMP(6))
WHERE created_at IS NULL;

UPDATE contact
SET
    created_at = COALESCE(created_at, UTC_TIMESTAMP(6)),
    updated_at = COALESCE(updated_at, created_at, UTC_TIMESTAMP(6))
WHERE created_at IS NULL OR updated_at IS NULL;

UPDATE target
SET
    created_at = COALESCE(created_at, UTC_TIMESTAMP(6)),
    updated_at = COALESCE(updated_at, created_at, UTC_TIMESTAMP(6))
WHERE created_at IS NULL OR updated_at IS NULL;

ALTER TABLE product
MODIFY COLUMN created_at DATETIME(6) NOT NULL,
MODIFY COLUMN updated_at DATETIME(6) NOT NULL;

ALTER TABLE category
MODIFY COLUMN created_at DATETIME(6) NOT NULL,
MODIFY COLUMN updated_at DATETIME(6) NOT NULL;

ALTER TABLE tag
MODIFY COLUMN created_at DATETIME(6) NOT NULL,
MODIFY COLUMN updated_at DATETIME(6) NOT NULL;

ALTER TABLE product_to_tag
MODIFY COLUMN created_at DATETIME(6) NOT NULL,
MODIFY COLUMN updated_at DATETIME(6) NOT NULL;

ALTER TABLE product_image
MODIFY COLUMN created_at DATETIME(6) NOT NULL,
MODIFY COLUMN updated_at DATETIME(6) NOT NULL;

ALTER TABLE sales_site
MODIFY COLUMN created_at DATETIME(6) NOT NULL,
MODIFY COLUMN updated_at DATETIME(6) NOT NULL;

ALTER TABLE site_detail
MODIFY COLUMN created_at DATETIME(6) NOT NULL,
MODIFY COLUMN updated_at DATETIME(6) NOT NULL;

ALTER TABLE creator
MODIFY COLUMN created_at DATETIME(6) NOT NULL,
MODIFY COLUMN updated_at DATETIME(6) NOT NULL;

ALTER TABLE skill_market
MODIFY COLUMN created_at DATETIME(6) NOT NULL,
MODIFY COLUMN updated_at DATETIME(6) NOT NULL;

ALTER TABLE sns
MODIFY COLUMN created_at DATETIME(6) NOT NULL,
MODIFY COLUMN updated_at DATETIME(6) NOT NULL;

ALTER TABLE user
MODIFY COLUMN created_at DATETIME(6) NOT NULL,
MODIFY COLUMN updated_at DATETIME(6) NOT NULL;

ALTER TABLE session
MODIFY COLUMN created_at DATETIME(6) NOT NULL;

ALTER TABLE contact
MODIFY COLUMN created_at DATETIME(6) NOT NULL,
MODIFY COLUMN updated_at DATETIME(6) NOT NULL;

ALTER TABLE target
MODIFY COLUMN created_at DATETIME(6) NOT NULL,
MODIFY COLUMN updated_at DATETIME(6) NOT NULL;
