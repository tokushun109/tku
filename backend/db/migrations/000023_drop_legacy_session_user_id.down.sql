-- 000023 の rollback では、user_uuid から user_id を復元して
-- 旧 index / foreign key を再作成する。

ALTER TABLE session
ADD COLUMN user_id INT NULL AFTER uuid;

UPDATE session s
INNER JOIN user u ON u.uuid = s.user_uuid
SET s.user_id = u.id
WHERE s.user_uuid IS NOT NULL;

ALTER TABLE session
ADD INDEX user_id (user_id);

ALTER TABLE session
ADD CONSTRAINT session_ibfk_1 FOREIGN KEY (user_id) REFERENCES user (id);
