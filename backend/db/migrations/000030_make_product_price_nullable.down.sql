-- NULL の price を補完してから NOT NULL 制約を復元する。
-- ドメインの価格下限（ProductPrice の最小値 1）に合わせて 1 で補完する。
-- 0 で補完するとロールバック後に NewProductPrice(0) が ErrInvalidPrice となり読み出しに失敗するため。
UPDATE product
SET price = 1
WHERE price IS NULL;

ALTER TABLE product
MODIFY COLUMN price INT NOT NULL DEFAULT 0;
