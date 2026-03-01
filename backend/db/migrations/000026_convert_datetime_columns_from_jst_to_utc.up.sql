-- Convert existing DATETIME values that were stored as JST wall-clock times
-- into UTC wall-clock times by subtracting 9 hours.
-- This migration assumes the application is switched to UTC immediately after.

START TRANSACTION;

UPDATE product
SET
    created_at = DATE_SUB(created_at, INTERVAL 9 HOUR),
    updated_at = DATE_SUB(updated_at, INTERVAL 9 HOUR),
    deleted_at = DATE_SUB(deleted_at, INTERVAL 9 HOUR);

UPDATE category
SET
    created_at = DATE_SUB(created_at, INTERVAL 9 HOUR),
    updated_at = DATE_SUB(updated_at, INTERVAL 9 HOUR),
    deleted_at = DATE_SUB(deleted_at, INTERVAL 9 HOUR);

UPDATE tag
SET
    created_at = DATE_SUB(created_at, INTERVAL 9 HOUR),
    updated_at = DATE_SUB(updated_at, INTERVAL 9 HOUR),
    deleted_at = DATE_SUB(deleted_at, INTERVAL 9 HOUR);

UPDATE product_to_tag
SET
    created_at = DATE_SUB(created_at, INTERVAL 9 HOUR),
    updated_at = DATE_SUB(updated_at, INTERVAL 9 HOUR),
    deleted_at = DATE_SUB(deleted_at, INTERVAL 9 HOUR);

UPDATE product_image
SET
    created_at = DATE_SUB(created_at, INTERVAL 9 HOUR),
    updated_at = DATE_SUB(updated_at, INTERVAL 9 HOUR),
    deleted_at = DATE_SUB(deleted_at, INTERVAL 9 HOUR);

UPDATE sales_site
SET
    created_at = DATE_SUB(created_at, INTERVAL 9 HOUR),
    updated_at = DATE_SUB(updated_at, INTERVAL 9 HOUR),
    deleted_at = DATE_SUB(deleted_at, INTERVAL 9 HOUR);

UPDATE site_detail
SET
    created_at = DATE_SUB(created_at, INTERVAL 9 HOUR),
    updated_at = DATE_SUB(updated_at, INTERVAL 9 HOUR);

UPDATE creator
SET
    created_at = DATE_SUB(created_at, INTERVAL 9 HOUR),
    updated_at = DATE_SUB(updated_at, INTERVAL 9 HOUR),
    deleted_at = DATE_SUB(deleted_at, INTERVAL 9 HOUR);

UPDATE skill_market
SET
    created_at = DATE_SUB(created_at, INTERVAL 9 HOUR),
    updated_at = DATE_SUB(updated_at, INTERVAL 9 HOUR),
    deleted_at = DATE_SUB(deleted_at, INTERVAL 9 HOUR);

UPDATE sns
SET
    created_at = DATE_SUB(created_at, INTERVAL 9 HOUR),
    updated_at = DATE_SUB(updated_at, INTERVAL 9 HOUR),
    deleted_at = DATE_SUB(deleted_at, INTERVAL 9 HOUR);

UPDATE user
SET
    created_at = DATE_SUB(created_at, INTERVAL 9 HOUR),
    updated_at = DATE_SUB(updated_at, INTERVAL 9 HOUR),
    deleted_at = DATE_SUB(deleted_at, INTERVAL 9 HOUR);

UPDATE session
SET
    created_at = DATE_SUB(created_at, INTERVAL 9 HOUR);

UPDATE contact
SET
    created_at = DATE_SUB(created_at, INTERVAL 9 HOUR),
    updated_at = DATE_SUB(updated_at, INTERVAL 9 HOUR),
    deleted_at = DATE_SUB(deleted_at, INTERVAL 9 HOUR);

UPDATE target
SET
    created_at = DATE_SUB(created_at, INTERVAL 9 HOUR),
    updated_at = DATE_SUB(updated_at, INTERVAL 9 HOUR),
    deleted_at = DATE_SUB(deleted_at, INTERVAL 9 HOUR);

COMMIT;
