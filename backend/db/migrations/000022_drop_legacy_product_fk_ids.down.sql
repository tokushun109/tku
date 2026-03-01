-- 000022 の rollback では、uuid 参照から旧 id 参照を復元する。
-- 先に列を戻し、uuid から id を逆引きしてバックフィルしたうえで、
-- 元の index / foreign key を再作成する。

-- product の category_id / target_id を復元する。
ALTER TABLE product
ADD COLUMN category_id INT NULL AFTER deleted_at;

ALTER TABLE product
ADD COLUMN target_id INT NULL AFTER category_uuid;

UPDATE product p
LEFT JOIN category c ON c.uuid = p.category_uuid
LEFT JOIN target t ON t.uuid = p.target_uuid
SET
    p.category_id = c.id,
    p.target_id = t.id
WHERE p.category_uuid IS NOT NULL OR p.target_uuid IS NOT NULL;

ALTER TABLE product
ADD INDEX category_id (category_id),
ADD INDEX target_id (target_id);

ALTER TABLE product
ADD CONSTRAINT product_ibfk_1 FOREIGN KEY (category_id) REFERENCES category (id),
ADD CONSTRAINT product_ibfk_2 FOREIGN KEY (target_id) REFERENCES target (id);

-- product_to_tag の product_id / tag_id を復元する。
ALTER TABLE product_to_tag
ADD COLUMN product_id INT NULL AFTER id;

ALTER TABLE product_to_tag
ADD COLUMN tag_id INT NULL AFTER product_uuid;

UPDATE product_to_tag ptt
INNER JOIN product p ON p.uuid = ptt.product_uuid
SET ptt.product_id = p.id
WHERE ptt.product_uuid IS NOT NULL;

UPDATE product_to_tag ptt
INNER JOIN tag t ON t.uuid = ptt.tag_uuid
SET ptt.tag_id = t.id
WHERE ptt.tag_uuid IS NOT NULL;

ALTER TABLE product_to_tag
ADD INDEX product_id (product_id),
ADD INDEX tag_id (tag_id);

ALTER TABLE product_to_tag
ADD CONSTRAINT product_to_tag_ibfk_1 FOREIGN KEY (product_id) REFERENCES product (id),
ADD CONSTRAINT product_to_tag_ibfk_2 FOREIGN KEY (tag_id) REFERENCES tag (id);

-- product_image の product_id を復元する。
ALTER TABLE product_image
ADD COLUMN product_id INT NULL AFTER deleted_at;

UPDATE product_image pi
INNER JOIN product p ON p.uuid = pi.product_uuid
SET pi.product_id = p.id
WHERE pi.product_uuid IS NOT NULL;

ALTER TABLE product_image
ADD INDEX product_id (product_id);

ALTER TABLE product_image
ADD CONSTRAINT product_image_ibfk_1 FOREIGN KEY (product_id) REFERENCES product (id);

-- site_detail の product_id / sales_site_id を復元する。
ALTER TABLE site_detail
ADD COLUMN product_id INT NULL AFTER detail_url;

ALTER TABLE site_detail
ADD COLUMN sales_site_id INT NULL AFTER product_uuid;

UPDATE site_detail sd
INNER JOIN product p ON p.uuid = sd.product_uuid
SET sd.product_id = p.id
WHERE sd.product_uuid IS NOT NULL;

UPDATE site_detail sd
INNER JOIN sales_site ss ON ss.uuid = sd.sales_site_uuid
SET sd.sales_site_id = ss.id
WHERE sd.sales_site_uuid IS NOT NULL;

ALTER TABLE site_detail
ADD INDEX product_id (product_id),
ADD INDEX sales_site_id (sales_site_id);

ALTER TABLE site_detail
ADD CONSTRAINT site_detail_ibfk_1 FOREIGN KEY (product_id) REFERENCES product (id),
ADD CONSTRAINT site_detail_ibfk_2 FOREIGN KEY (sales_site_id) REFERENCES sales_site (id);
