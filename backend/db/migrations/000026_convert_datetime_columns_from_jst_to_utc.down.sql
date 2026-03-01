-- Revert the JST->UTC DATETIME conversion by adding 9 hours back.
-- Use only before new UTC-based writes are introduced, or together with a full rollback.

START TRANSACTION;

UPDATE product
SET
    created_at = DATE_ADD(created_at, INTERVAL 9 HOUR),
    updated_at = DATE_ADD(updated_at, INTERVAL 9 HOUR),
    deleted_at = DATE_ADD(deleted_at, INTERVAL 9 HOUR);

UPDATE category
SET
    created_at = DATE_ADD(created_at, INTERVAL 9 HOUR),
    updated_at = DATE_ADD(updated_at, INTERVAL 9 HOUR),
    deleted_at = DATE_ADD(deleted_at, INTERVAL 9 HOUR);

UPDATE tag
SET
    created_at = DATE_ADD(created_at, INTERVAL 9 HOUR),
    updated_at = DATE_ADD(updated_at, INTERVAL 9 HOUR),
    deleted_at = DATE_ADD(deleted_at, INTERVAL 9 HOUR);

UPDATE product_to_tag
SET
    created_at = DATE_ADD(created_at, INTERVAL 9 HOUR),
    updated_at = DATE_ADD(updated_at, INTERVAL 9 HOUR),
    deleted_at = DATE_ADD(deleted_at, INTERVAL 9 HOUR);

UPDATE product_image
SET
    created_at = DATE_ADD(created_at, INTERVAL 9 HOUR),
    updated_at = DATE_ADD(updated_at, INTERVAL 9 HOUR),
    deleted_at = DATE_ADD(deleted_at, INTERVAL 9 HOUR);

UPDATE sales_site
SET
    created_at = DATE_ADD(created_at, INTERVAL 9 HOUR),
    updated_at = DATE_ADD(updated_at, INTERVAL 9 HOUR),
    deleted_at = DATE_ADD(deleted_at, INTERVAL 9 HOUR);

UPDATE site_detail
SET
    created_at = DATE_ADD(created_at, INTERVAL 9 HOUR),
    updated_at = DATE_ADD(updated_at, INTERVAL 9 HOUR);

UPDATE creator
SET
    created_at = DATE_ADD(created_at, INTERVAL 9 HOUR),
    updated_at = DATE_ADD(updated_at, INTERVAL 9 HOUR),
    deleted_at = DATE_ADD(deleted_at, INTERVAL 9 HOUR);

UPDATE skill_market
SET
    created_at = DATE_ADD(created_at, INTERVAL 9 HOUR),
    updated_at = DATE_ADD(updated_at, INTERVAL 9 HOUR),
    deleted_at = DATE_ADD(deleted_at, INTERVAL 9 HOUR);

UPDATE sns
SET
    created_at = DATE_ADD(created_at, INTERVAL 9 HOUR),
    updated_at = DATE_ADD(updated_at, INTERVAL 9 HOUR),
    deleted_at = DATE_ADD(deleted_at, INTERVAL 9 HOUR);

UPDATE user
SET
    created_at = DATE_ADD(created_at, INTERVAL 9 HOUR),
    updated_at = DATE_ADD(updated_at, INTERVAL 9 HOUR),
    deleted_at = DATE_ADD(deleted_at, INTERVAL 9 HOUR);

UPDATE session
SET
    created_at = DATE_ADD(created_at, INTERVAL 9 HOUR);

UPDATE contact
SET
    created_at = DATE_ADD(created_at, INTERVAL 9 HOUR),
    updated_at = DATE_ADD(updated_at, INTERVAL 9 HOUR),
    deleted_at = DATE_ADD(deleted_at, INTERVAL 9 HOUR);

UPDATE target
SET
    created_at = DATE_ADD(created_at, INTERVAL 9 HOUR),
    updated_at = DATE_ADD(updated_at, INTERVAL 9 HOUR),
    deleted_at = DATE_ADD(deleted_at, INTERVAL 9 HOUR);

COMMIT;
