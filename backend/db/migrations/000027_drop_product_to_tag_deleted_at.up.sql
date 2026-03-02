-- product_to_tag はアプリケーション上ですでに物理削除運用になっているため、
-- 旧運用で論理削除済みになっている行を掃除してから deleted_at を削除する。
DELETE FROM product_to_tag
WHERE deleted_at IS NOT NULL;

ALTER TABLE product_to_tag
DROP COLUMN deleted_at;
