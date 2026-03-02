-- 高優先度のスキーマ改善:
-- 1. category / target / tag / sales_site の「有効レコード名」重複を解消し、一意制約を付与する
-- 2. session を user_uuid 単位で 1 レコードに制約する

-- category の同名有効レコードは最小 id を正とし、product 参照を寄せてから重複行を論理削除する。
UPDATE product p
INNER JOIN (
    SELECT
        dup.uuid AS duplicate_uuid,
        keep.uuid AS canonical_uuid
    FROM category dup
    INNER JOIN category keep
        ON keep.name = dup.name
        AND keep.deleted_at IS NULL
        AND keep.id < dup.id
    LEFT JOIN category earlier
        ON earlier.name = dup.name
        AND earlier.deleted_at IS NULL
        AND earlier.id < keep.id
    WHERE
        dup.deleted_at IS NULL
        AND earlier.id IS NULL
) mapping
    ON mapping.duplicate_uuid = p.category_uuid
SET p.category_uuid = mapping.canonical_uuid;

UPDATE category dup
INNER JOIN category keep
    ON keep.name = dup.name
    AND keep.deleted_at IS NULL
    AND keep.id < dup.id
SET
    dup.deleted_at = COALESCE(dup.deleted_at, UTC_TIMESTAMP()),
    dup.updated_at = UTC_TIMESTAMP()
WHERE dup.deleted_at IS NULL;

ALTER TABLE category
ADD COLUMN active_name VARCHAR(30)
GENERATED ALWAYS AS (
    CASE
        WHEN deleted_at IS NULL THEN name
        ELSE NULL
    END
) STORED;

CREATE UNIQUE INDEX idx_category_active_name ON category (active_name);

-- target も category と同様に、product 参照を正規化してから有効名の一意制約を付与する。
UPDATE product p
INNER JOIN (
    SELECT
        dup.uuid AS duplicate_uuid,
        keep.uuid AS canonical_uuid
    FROM target dup
    INNER JOIN target keep
        ON keep.name = dup.name
        AND keep.deleted_at IS NULL
        AND keep.id < dup.id
    LEFT JOIN target earlier
        ON earlier.name = dup.name
        AND earlier.deleted_at IS NULL
        AND earlier.id < keep.id
    WHERE
        dup.deleted_at IS NULL
        AND earlier.id IS NULL
) mapping
    ON mapping.duplicate_uuid = p.target_uuid
SET p.target_uuid = mapping.canonical_uuid;

UPDATE target dup
INNER JOIN target keep
    ON keep.name = dup.name
    AND keep.deleted_at IS NULL
    AND keep.id < dup.id
SET
    dup.deleted_at = COALESCE(dup.deleted_at, UTC_TIMESTAMP()),
    dup.updated_at = UTC_TIMESTAMP()
WHERE dup.deleted_at IS NULL;

ALTER TABLE target
ADD COLUMN active_name VARCHAR(30)
GENERATED ALWAYS AS (
    CASE
        WHEN deleted_at IS NULL THEN name
        ELSE NULL
    END
) STORED;

CREATE UNIQUE INDEX idx_target_active_name ON target (active_name);

-- tag は product_to_tag の参照を正規化する。
-- 先に「正規化後に重複する中間テーブル行」を削除し、その後 tag_uuid を canonical へ寄せる。
DELETE ptt FROM product_to_tag ptt
INNER JOIN (
    SELECT
        dup.uuid AS duplicate_uuid,
        keep.uuid AS canonical_uuid
    FROM tag dup
    INNER JOIN tag keep
        ON keep.name = dup.name
        AND keep.deleted_at IS NULL
        AND keep.id < dup.id
    LEFT JOIN tag earlier
        ON earlier.name = dup.name
        AND earlier.deleted_at IS NULL
        AND earlier.id < keep.id
    WHERE
        dup.deleted_at IS NULL
        AND earlier.id IS NULL
) mapping
    ON mapping.duplicate_uuid = ptt.tag_uuid
INNER JOIN product_to_tag canonical_ptt
    ON canonical_ptt.product_uuid = ptt.product_uuid
    AND canonical_ptt.tag_uuid = mapping.canonical_uuid;

UPDATE product_to_tag ptt
INNER JOIN (
    SELECT
        dup.uuid AS duplicate_uuid,
        keep.uuid AS canonical_uuid
    FROM tag dup
    INNER JOIN tag keep
        ON keep.name = dup.name
        AND keep.deleted_at IS NULL
        AND keep.id < dup.id
    LEFT JOIN tag earlier
        ON earlier.name = dup.name
        AND earlier.deleted_at IS NULL
        AND earlier.id < keep.id
    WHERE
        dup.deleted_at IS NULL
        AND earlier.id IS NULL
) mapping
    ON mapping.duplicate_uuid = ptt.tag_uuid
SET ptt.tag_uuid = mapping.canonical_uuid;

UPDATE tag dup
INNER JOIN tag keep
    ON keep.name = dup.name
    AND keep.deleted_at IS NULL
    AND keep.id < dup.id
SET
    dup.deleted_at = COALESCE(dup.deleted_at, UTC_TIMESTAMP()),
    dup.updated_at = UTC_TIMESTAMP()
WHERE dup.deleted_at IS NULL;

ALTER TABLE tag
ADD COLUMN active_name VARCHAR(30)
GENERATED ALWAYS AS (
    CASE
        WHEN deleted_at IS NULL THEN name
        ELSE NULL
    END
) STORED;

CREATE UNIQUE INDEX idx_tag_active_name ON tag (active_name);

-- sales_site は site_detail の参照を正規化する。
-- site_detail は (product_uuid, sales_site_uuid) に一意制約があるため、衝突行を先に削除する。
DELETE sd FROM site_detail sd
INNER JOIN (
    SELECT
        dup.uuid AS duplicate_uuid,
        keep.uuid AS canonical_uuid
    FROM sales_site dup
    INNER JOIN sales_site keep
        ON keep.name = dup.name
        AND keep.deleted_at IS NULL
        AND keep.id < dup.id
    LEFT JOIN sales_site earlier
        ON earlier.name = dup.name
        AND earlier.deleted_at IS NULL
        AND earlier.id < keep.id
    WHERE
        dup.deleted_at IS NULL
        AND earlier.id IS NULL
) mapping
    ON mapping.duplicate_uuid = sd.sales_site_uuid
INNER JOIN site_detail canonical_sd
    ON canonical_sd.product_uuid = sd.product_uuid
    AND canonical_sd.sales_site_uuid = mapping.canonical_uuid;

UPDATE site_detail sd
INNER JOIN (
    SELECT
        dup.uuid AS duplicate_uuid,
        keep.uuid AS canonical_uuid
    FROM sales_site dup
    INNER JOIN sales_site keep
        ON keep.name = dup.name
        AND keep.deleted_at IS NULL
        AND keep.id < dup.id
    LEFT JOIN sales_site earlier
        ON earlier.name = dup.name
        AND earlier.deleted_at IS NULL
        AND earlier.id < keep.id
    WHERE
        dup.deleted_at IS NULL
        AND earlier.id IS NULL
) mapping
    ON mapping.duplicate_uuid = sd.sales_site_uuid
SET sd.sales_site_uuid = mapping.canonical_uuid;

UPDATE sales_site dup
INNER JOIN sales_site keep
    ON keep.name = dup.name
    AND keep.deleted_at IS NULL
    AND keep.id < dup.id
SET
    dup.deleted_at = COALESCE(dup.deleted_at, UTC_TIMESTAMP()),
    dup.updated_at = UTC_TIMESTAMP()
WHERE dup.deleted_at IS NULL;

ALTER TABLE sales_site
ADD COLUMN active_name VARCHAR(30)
GENERATED ALWAYS AS (
    CASE
        WHEN deleted_at IS NULL THEN name
        ELSE NULL
    END
) STORED;

CREATE UNIQUE INDEX idx_sales_site_active_name ON sales_site (active_name);

-- session は user_uuid ごとに最新 1 件を残し、それ以外を削除してから UNIQUE 化する。
DELETE s1 FROM session s1
INNER JOIN session s2
WHERE
    s1.user_uuid = s2.user_uuid
    AND (
        COALESCE(s1.created_at, '1000-01-01 00:00:00') < COALESCE(s2.created_at, '1000-01-01 00:00:00')
        OR (
            COALESCE(s1.created_at, '1000-01-01 00:00:00') = COALESCE(s2.created_at, '1000-01-01 00:00:00')
            AND s1.id < s2.id
        )
    );

ALTER TABLE session
DROP FOREIGN KEY fk_session_user_uuid;

DROP INDEX idx_session_user_uuid ON session;

CREATE UNIQUE INDEX idx_session_user_uuid ON session (user_uuid);

ALTER TABLE session
ADD CONSTRAINT fk_session_user_uuid
FOREIGN KEY (user_uuid) REFERENCES user (uuid);
