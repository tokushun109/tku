-- rollback では site_detail の deleted_at 列定義のみを戻す。
-- up で物理削除した行自体は復元しない。
ALTER TABLE site_detail
ADD COLUMN deleted_at DATETIME NULL AFTER updated_at;
