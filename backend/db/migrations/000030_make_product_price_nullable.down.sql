-- NULL の price を 0 で補完してから NOT NULL 制約を復元する。
UPDATE product
SET price = 0
WHERE price IS NULL;

ALTER TABLE product
MODIFY COLUMN price INT NOT NULL DEFAULT 0;
