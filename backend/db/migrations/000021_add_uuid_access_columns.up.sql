-- UUID ベースアクセスへ移行するための追加 migration。
-- 既存の id カラムは当面残しつつ、外部公開・関連付けで使う列を uuid 側へ寄せる。
-- 基本方針は以下の通り:
-- 1. まず新しい uuid 系カラムを NULL 許容で追加する
-- 2. 既存の id 系外部キーから値をバックフィルする
-- 3. インデックス / UNIQUE / FK を追加して、以降の UUID 参照を成立させる
-- この migration 自体では旧 id 列は削除しない。アプリケーション切替後に段階的に整理する想定。

-- creator はこれまで id のみだったため、外部識別用の uuid を新規付与する。
-- 既存レコードにも一意な値が必要なので、一度 NULL 許容で追加してから採番する。
ALTER TABLE creator
ADD COLUMN uuid VARCHAR(36) NULL AFTER id;

-- 既存 creator レコードに lowercase の UUID を採番する。
-- アプリケーション側では lowercase 前提で比較するため、ここでも小文字に揃える。
UPDATE creator
SET uuid = LOWER(UUID())
WHERE uuid IS NULL;

-- 全件バックフィル後、NOT NULL に変更して必須カラム化する。
ALTER TABLE creator
MODIFY COLUMN uuid VARCHAR(36) NOT NULL;

-- creator.uuid を外部識別子として利用できるよう、UNIQUE 制約相当の一意インデックスを追加する。
CREATE UNIQUE INDEX idx_creator_uuid ON creator (uuid);

-- contact も creator と同様に、これまで id のみだったため uuid を新規付与する。
ALTER TABLE contact
ADD COLUMN uuid VARCHAR(36) NULL AFTER id;

-- 既存 contact レコードに lowercase の UUID を採番する。
UPDATE contact
SET uuid = LOWER(UUID())
WHERE uuid IS NULL;

-- バックフィル完了後に NOT NULL 化する。
ALTER TABLE contact
MODIFY COLUMN uuid VARCHAR(36) NOT NULL;

-- contact.uuid を API レスポンスや外部参照の主識別子として使えるよう一意インデックスを付与する。
CREATE UNIQUE INDEX idx_contact_uuid ON contact (uuid);

-- product は category / target の関連先を id ではなく uuid で保持できるようにする。
-- 移行期間中に旧データを扱えるよう、追加時点では NULL 許容にしている。
ALTER TABLE product
ADD COLUMN category_uuid VARCHAR(36) NULL AFTER category_id,
ADD COLUMN target_uuid VARCHAR(36) NULL AFTER target_id;

-- 既存の category_id / target_id から、それぞれ対応する uuid をバックフィルする。
-- LEFT JOIN にしているのは、片方だけ NULL のレコードもそのまま扱えるようにするため。
UPDATE product p
LEFT JOIN category c ON c.id = p.category_id
LEFT JOIN target t ON t.id = p.target_id
SET
    p.category_uuid = c.uuid,
    p.target_uuid = t.uuid
WHERE p.category_id IS NOT NULL OR p.target_id IS NOT NULL;

-- UUID 列での検索・JOIN を前提に、個別インデックスを追加する。
CREATE INDEX idx_product_category_uuid ON product (category_uuid);
CREATE INDEX idx_product_target_uuid ON product (target_uuid);

-- 以後は category.uuid / target.uuid を参照元にできるよう FK を追加する。
-- 旧 category_id / target_id は当面残し、アプリ切替後に段階的に整理する。
ALTER TABLE product
ADD CONSTRAINT fk_product_category_uuid
FOREIGN KEY (category_uuid) REFERENCES category (uuid);

ALTER TABLE product
ADD CONSTRAINT fk_product_target_uuid
FOREIGN KEY (target_uuid) REFERENCES target (uuid);

-- product_to_tag は中間テーブル。今後は product_uuid / tag_uuid で関連を持てるようにする。
-- 旧 id 系カラムを NULL 許容へ緩めておくことで、移行期間中に両方の列を共存させる。
ALTER TABLE product_to_tag
MODIFY COLUMN product_id INT NULL,
MODIFY COLUMN tag_id INT NULL,
ADD COLUMN product_uuid VARCHAR(36) NULL AFTER product_id,
ADD COLUMN tag_uuid VARCHAR(36) NULL AFTER tag_id;

-- 既存の product_id / tag_id から中間テーブルの UUID 参照をバックフィルする。
UPDATE product_to_tag ptt
INNER JOIN product p ON p.id = ptt.product_id
INNER JOIN tag t ON t.id = ptt.tag_id
SET
    ptt.product_uuid = p.uuid,
    ptt.tag_uuid = t.uuid
WHERE ptt.product_id IS NOT NULL AND ptt.tag_id IS NOT NULL;

-- product_uuid / tag_uuid 単体の検索用インデックスと、
-- 同じ組み合わせが重複しないよう複合 UNIQUE インデックスを追加する。
CREATE INDEX idx_product_to_tag_product_uuid ON product_to_tag (product_uuid);
CREATE INDEX idx_product_to_tag_tag_uuid ON product_to_tag (tag_uuid);
CREATE UNIQUE INDEX idx_product_to_tag_product_uuid_tag_uuid ON product_to_tag (product_uuid, tag_uuid);

-- 今後の関連整合性を担保するため、uuid ベースの FK を追加する。
ALTER TABLE product_to_tag
ADD CONSTRAINT fk_product_to_tag_product_uuid
FOREIGN KEY (product_uuid) REFERENCES product (uuid);

ALTER TABLE product_to_tag
ADD CONSTRAINT fk_product_to_tag_tag_uuid
FOREIGN KEY (tag_uuid) REFERENCES tag (uuid);

-- product_image も product_id 参照から product_uuid 参照へ移行できるようにする。
ALTER TABLE product_image
ADD COLUMN product_uuid VARCHAR(36) NULL AFTER product_id;

-- 既存 product_id から product_uuid をバックフィルする。
UPDATE product_image pi
INNER JOIN product p ON p.id = pi.product_id
SET pi.product_uuid = p.uuid
WHERE pi.product_id IS NOT NULL;

-- product_uuid での参照性能を確保するためインデックスを追加する。
CREATE INDEX idx_product_image_product_uuid ON product_image (product_uuid);

-- product.uuid に対する FK を追加し、以後の整合性を担保する。
ALTER TABLE product_image
ADD CONSTRAINT fk_product_image_product_uuid
FOREIGN KEY (product_uuid) REFERENCES product (uuid);

-- site_detail は product / sales_site の両方を参照しているため、両方の UUID 列を追加する。
-- 旧 id 列は移行期間中の互換のため NULL 許容へ変更して残す。
ALTER TABLE site_detail
MODIFY COLUMN product_id INT NULL,
MODIFY COLUMN sales_site_id INT NULL,
ADD COLUMN product_uuid VARCHAR(36) NULL AFTER product_id,
ADD COLUMN sales_site_uuid VARCHAR(36) NULL AFTER sales_site_id;

-- 既存の id 参照から product_uuid / sales_site_uuid をバックフィルする。
UPDATE site_detail sd
INNER JOIN product p ON p.id = sd.product_id
INNER JOIN sales_site ss ON ss.id = sd.sales_site_id
SET
    sd.product_uuid = p.uuid,
    sd.sales_site_uuid = ss.uuid
WHERE sd.product_id IS NOT NULL AND sd.sales_site_id IS NOT NULL;

-- UUID 参照での検索・JOIN 用インデックスを追加する。
CREATE INDEX idx_site_detail_product_uuid ON site_detail (product_uuid);
CREATE INDEX idx_site_detail_sales_site_uuid ON site_detail (sales_site_uuid);

-- 新しい UUID 列に対して FK を追加し、参照整合性を担保する。
ALTER TABLE site_detail
ADD CONSTRAINT fk_site_detail_product_uuid
FOREIGN KEY (product_uuid) REFERENCES product (uuid);

ALTER TABLE site_detail
ADD CONSTRAINT fk_site_detail_sales_site_uuid
FOREIGN KEY (sales_site_uuid) REFERENCES sales_site (uuid);

-- session は user_id 参照から user_uuid 参照へ移行する。
-- 認証系は段階切替を想定しているため、旧 user_id は残したまま NULL 許容へ変更する。
ALTER TABLE session
MODIFY COLUMN user_id INT NULL,
ADD COLUMN user_uuid VARCHAR(36) NULL AFTER user_id;

-- 既存 session の user_id から user_uuid をバックフィルする。
UPDATE session s
INNER JOIN user u ON u.id = s.user_id
SET s.user_uuid = u.uuid
WHERE s.user_id IS NOT NULL;

-- user_uuid でのセッション解決に備えてインデックスを追加する。
CREATE INDEX idx_session_user_uuid ON session (user_uuid);

-- 以後は user.uuid を参照元にできるよう FK を追加する。
ALTER TABLE session
ADD CONSTRAINT fk_session_user_uuid
FOREIGN KEY (user_uuid) REFERENCES user (uuid);
