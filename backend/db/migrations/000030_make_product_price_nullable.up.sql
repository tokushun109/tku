-- 価格を任意項目化する:
-- 販売サイトごとに価格が異なるため、価格を未設定にできるよう NULL 許容に変更する。
ALTER TABLE product
MODIFY COLUMN price INT NULL;

-- 旧スキーマの DEFAULT 0 により「未設定」相当のまま残っている 0 円レコードを
-- NULL（未設定）へ正規化する。ドメインの価格下限は 1 のため 0 は有効価格ではない。
UPDATE product
SET price = NULL
WHERE price = 0;
