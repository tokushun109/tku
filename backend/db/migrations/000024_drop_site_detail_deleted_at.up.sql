-- site_detail はアプリケーション上ですでに物理削除運用になっているため、
-- 旧運用で論理削除済みになっている行を掃除してから deleted_at を削除する。
DELETE FROM site_detail
WHERE deleted_at IS NOT NULL;

ALTER TABLE site_detail
DROP COLUMN deleted_at;
