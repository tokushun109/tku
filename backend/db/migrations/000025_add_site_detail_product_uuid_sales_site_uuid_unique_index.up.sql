-- 既存データに (product_uuid, sales_site_uuid) の組み合わせで重複が存在する場合に備え、
-- CREATE UNIQUE INDEX の前に重複を削除する。id が最も小さいものを残す。
DELETE t1 FROM site_detail t1
INNER JOIN site_detail t2
WHERE
    t1.id > t2.id AND
    t1.product_uuid = t2.product_uuid AND
    t1.sales_site_uuid = t2.sales_site_uuid;

-- site_detail は 1 商品につき 1 販売サイト 1 件に制約する。
CREATE UNIQUE INDEX idx_site_detail_product_uuid_sales_site_uuid
ON site_detail (product_uuid, sales_site_uuid);
