-- site_detail は 1 商品につき 1 販売サイト 1 件に制約する。
CREATE UNIQUE INDEX idx_site_detail_product_uuid_sales_site_uuid
ON site_detail (product_uuid, sales_site_uuid);
