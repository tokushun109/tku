-- rollback では product_to_tag の deleted_at 列定義のみを戻す。
-- up で物理削除した行自体は復元しない。
ALTER TABLE product_to_tag
ADD COLUMN deleted_at DATETIME NULL AFTER updated_at;
