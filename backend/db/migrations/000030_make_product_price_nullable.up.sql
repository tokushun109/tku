-- 価格を任意項目化する:
-- 販売サイトごとに価格が異なるため、価格を未設定にできるよう NULL 許容に変更する。
ALTER TABLE product
MODIFY COLUMN price INT NULL;
